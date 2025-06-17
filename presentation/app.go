package presentation

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"karedoro/application"
	"karedoro/domain"
)

type App struct {
	sessionService      *application.SessionService
	audioService        *application.AudioService
	notificationService *application.NotificationService
	configService       *application.ConfigService
	
	currentScreen   Screen
	isFullscreen    bool
	buttonManager   *ButtonManager
}

type Screen int

const (
	MainScreen Screen = iota
	FullscreenOverlay
)

type Button struct {
	X, Y, W, H int
	Text       string
	Action     func()
	Hovered    bool
}

func NewApp() *App {
	app := &App{
		sessionService:      application.NewSessionService(),
		audioService:        application.NewAudioService(),
		notificationService: application.NewNotificationService(),
		configService:       application.NewConfigService(),
		currentScreen:       MainScreen,
		isFullscreen:        false,
		buttonManager:       NewButtonManager(),
	}
	
	app.setupEventCallbacks()
	app.setupButtons()
	
	return app
}

func (a *App) setupEventCallbacks() {
	a.sessionService.AddEventCallback(domain.EventWorkSessionStart, func() {
		a.audioService.PlayStartSound()
		a.notificationService.ShowWorkSessionStart()
		a.currentScreen = MainScreen
		if a.isFullscreen {
			ebiten.SetFullscreen(false)
			a.isFullscreen = false
		}
	})
	
	a.sessionService.AddEventCallback(domain.EventBreakSessionStart, func() {
		a.audioService.PlayStartSound()
		a.notificationService.ShowBreakSessionStart()
		a.currentScreen = MainScreen
		if a.isFullscreen {
			ebiten.SetFullscreen(false)
			a.isFullscreen = false
		}
	})
	
	a.sessionService.AddEventCallback(domain.EventWorkSessionEnd, func() {
		a.audioService.PlayEndSound()
		a.notificationService.ShowWorkSessionEnd()
		a.currentScreen = FullscreenOverlay
		ebiten.SetFullscreen(true)
		a.isFullscreen = true
		screenWidth, screenHeight := ebiten.WindowSize()
		a.buttonManager.SetupEndOfWorkButtons(screenWidth, screenHeight, a.sessionService)
	})
	
	a.sessionService.AddEventCallback(domain.EventBreakSessionEnd, func() {
		a.audioService.PlayEndSound()
		a.notificationService.ShowBreakSessionEnd()
		a.currentScreen = FullscreenOverlay
		ebiten.SetFullscreen(true)
		a.isFullscreen = true
		screenWidth, screenHeight := ebiten.WindowSize()
		a.buttonManager.SetupEndOfBreakButtons(screenWidth, screenHeight, a.sessionService)
	})
	
	a.sessionService.AddEventCallback(domain.EventWarning, func() {
		a.audioService.PlayWarningSound()
		a.notificationService.ShowWarning()
	})
	
	a.sessionService.AddEventCallback(domain.EventSessionPause, func() {
		a.audioService.PlayBeep(400, 100*time.Millisecond)
		a.notificationService.ShowSessionPaused()
	})
	
	a.sessionService.AddEventCallback(domain.EventSessionResume, func() {
		a.audioService.PlayBeep(600, 100*time.Millisecond)
		a.notificationService.ShowSessionResumed()
	})
}

func (a *App) setupButtons() {
	screenWidth, screenHeight := ebiten.WindowSize()
	a.buttonManager.SetupMainButtons(screenWidth, screenHeight, a.sessionService)
}


func (a *App) Update() error {
	a.sessionService.Update()
	
	session := a.sessionService.GetSession()
	
	switch session.GetState() {
	case domain.WorkSession, domain.BreakSession:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			if session.IsSessionPaused() {
				a.sessionService.ResumeSession()
			} else {
				a.sessionService.PauseSession()
			}
		}
	}
	
	a.updateButtonPositions()
	a.buttonManager.UpdateButtons()
	
	return nil
}

func (a *App) updateButtonPositions() {
	screenWidth, screenHeight := ebiten.WindowSize()
	a.buttonManager.UpdateButtonPositions(screenWidth, screenHeight)
}


func (a *App) Draw(screen *ebiten.Image) {
	screen.Fill(BackgroundColor)
	
	switch a.currentScreen {
	case MainScreen:
		a.drawMainScreen(screen)
	case FullscreenOverlay:
		a.drawFullscreenOverlay(screen)
	}
}

func (a *App) drawMainScreen(screen *ebiten.Image) {
	session := a.sessionService.GetSession()
	
	switch session.GetState() {
	case domain.WorkSession:
		a.drawWorkSession(screen)
	case domain.BreakSession:
		a.drawBreakSession(screen)
	case domain.Idle:
		a.drawIdleScreen(screen)
	}
}

func (a *App) drawWorkSession(screen *ebiten.Image) {
	session := a.sessionService.GetSession()
	remaining := session.GetTimeRemaining()
	screenWidth, screenHeight := ebiten.WindowSize()
	
	screen.Fill(WorkSessionColor)
	
	timerText := fmt.Sprintf("%02d:%02d", int(remaining.Minutes()), int(remaining.Seconds())%60)
	// Draw timer background for better visibility
	timerX := screenWidth/2 - TimerOffsetX
	timerY := screenHeight/2 - TimerOffsetY
	drawRect(screen, timerX, timerY, TimerBoxWidth, TimerBoxHeight, BlackShadow)
	ebitenutil.DebugPrintAt(screen, timerText, screenWidth/2-TimerOffsetX+10, screenHeight/2-TimerOffsetY+10)
	
	// Draw progress bar
	a.drawProgressBar(screen, session.GetProgress(), screenWidth, screenHeight)
	
	if session.IsSessionPaused() {
		ebitenutil.DebugPrintAt(screen, PausedText, screenWidth/2-len(PausedText)*TextCharWidth, screenHeight/2-TextLineHeight)
		ebitenutil.DebugPrintAt(screen, ResumeInstructionText, screenWidth/2-len(ResumeInstructionText)*TextCharWidth, screenHeight/2-20)
	} else {
		ebitenutil.DebugPrintAt(screen, WorkingText, screenWidth/2-len(WorkingText)*TextCharWidth, screenHeight/2-TextLineHeight)
		ebitenutil.DebugPrintAt(screen, PauseInstructionText, screenWidth/2-len(PauseInstructionText)*TextCharWidth, screenHeight/2-20)
	}
}

func (a *App) drawBreakSession(screen *ebiten.Image) {
	session := a.sessionService.GetSession()
	remaining := session.GetTimeRemaining()
	screenWidth, screenHeight := ebiten.WindowSize()
	
	screen.Fill(BreakSessionColor)
	
	timerText := fmt.Sprintf("%02d:%02d", int(remaining.Minutes()), int(remaining.Seconds())%60)
	// Draw timer background for better visibility
	timerX := screenWidth/2 - TimerOffsetX
	timerY := screenHeight/2 - TimerOffsetY
	drawRect(screen, timerX, timerY, TimerBoxWidth, TimerBoxHeight, BlackShadow)
	ebitenutil.DebugPrintAt(screen, timerText, screenWidth/2-TimerOffsetX+10, screenHeight/2-TimerOffsetY+10)
	
	// Draw progress bar
	a.drawProgressBar(screen, session.GetProgress(), screenWidth, screenHeight)
	
	if session.IsSessionPaused() {
		ebitenutil.DebugPrintAt(screen, PausedText, screenWidth/2-len(PausedText)*TextCharWidth, screenHeight/2-TextLineHeight)
		ebitenutil.DebugPrintAt(screen, ResumeInstructionText, screenWidth/2-len(ResumeInstructionText)*TextCharWidth, screenHeight/2-20)
	} else {
		ebitenutil.DebugPrintAt(screen, BreakText, screenWidth/2-len(BreakText)*TextCharWidth, screenHeight/2-TextLineHeight)
		ebitenutil.DebugPrintAt(screen, PauseInstructionText, screenWidth/2-len(PauseInstructionText)*TextCharWidth, screenHeight/2-20)
	}
}

func (a *App) drawIdleScreen(screen *ebiten.Image) {
	screenWidth, screenHeight := ebiten.WindowSize()
	ebitenutil.DebugPrintAt(screen, IdleScreenMessage, screenWidth/2-len(IdleScreenMessage)*TextCharWidth, screenHeight/2-IdleMessageOffset)
	
	a.buttonManager.DrawButtons(screen)
}

func (a *App) drawFullscreenOverlay(screen *ebiten.Image) {
	session := a.sessionService.GetSession()
	screenWidth, screenHeight := ebiten.WindowSize()
	
	// 強制的な赤い背景で注意を引く
	screen.Fill(ForceRedBackground)
	
	var message string
	if session.GetSessionType() == domain.Work {
		message = WorkSessionEndMessage
	} else {
		message = BreakSessionEndMessage
	}
	
	// メッセージを大きく強調表示
	messageX := screenWidth/2 - len(message)*TextCharWidthLg
	messageY := screenHeight/3
	
	// 背景の強調ボックス
	boxWidth := len(message)*8 + MessageBoxPadding
	boxHeight := MessageBoxHeight
	boxX := screenWidth/2 - boxWidth/2
	boxY := messageY - 20
	drawRect(screen, boxX, boxY, boxWidth, boxHeight, ForceYellowBox)
	drawBorder(screen, boxX, boxY, boxWidth, boxHeight, WhiteBorder, MessageBoxBorderWidth)
	
	ebitenutil.DebugPrintAt(screen, message, messageX, messageY)
	
	// 警告メッセージを追加
	warningMsg := "YOU CANNOT CONTINUE UNTIL YOU CHOOSE!"
	warningX := screenWidth/2 - len(warningMsg)*TextCharWidth
	warningY := screenHeight/2 - TextLineHeight
	ebitenutil.DebugPrintAt(screen, warningMsg, warningX, warningY)
	
	a.buttonManager.DrawButtons(screen)
}


func (a *App) drawProgressBar(screen *ebiten.Image, progress float64, screenWidth, screenHeight int) {
	barX := screenWidth/2 - ProgressBarWidth/2
	barY := screenHeight/2 + ProgressBarOffsetY
	
	// Draw progress bar background
	drawRect(screen, barX, barY, ProgressBarWidth, ProgressBarHeight, ProgressBarBg)
	
	// Draw progress bar fill
	progressWidth := int(float64(ProgressBarWidth) * progress)
	if progressWidth > 0 {
		drawRect(screen, barX, barY, progressWidth, ProgressBarHeight, ProgressBarFill)
	}
	
	// Draw progress bar border
	drawBorder(screen, barX, barY, ProgressBarWidth, ProgressBarHeight, ProgressBarBorder, 1)
}

func (a *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// Allow dynamic resizing but set minimum size
	if outsideWidth < MinWindowWidth {
		outsideWidth = MinWindowWidth
	}
	if outsideHeight < MinWindowHeight {
		outsideHeight = MinWindowHeight
	}
	
	return outsideWidth, outsideHeight
}

func (a *App) Run() error {
	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowTitle(WindowTitle)
	ebiten.SetWindowClosingHandled(true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	
	if !a.audioService.WaitForReady(5 * time.Second) {
		fmt.Println("Warning: Audio system initialization timeout")
	}
	
	return ebiten.RunGame(a)
}

func (a *App) IsDrawingSkipped() bool {
	return false
}