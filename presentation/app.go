package presentation

import (
	"fmt"
	"image/color"
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
	
	currentScreen Screen
	isFullscreen  bool
	
	buttons []Button
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
		buttons:            make([]Button, 0),
	}
	
	app.setupEventCallbacks()
	app.setupButtons()
	
	return app
}

func (a *App) setupEventCallbacks() {
	a.sessionService.AddEventCallback("work_session_start", func() {
		a.audioService.PlayStartSound()
		a.notificationService.ShowWorkSessionStart()
		a.currentScreen = MainScreen
		if a.isFullscreen {
			ebiten.SetFullscreen(false)
			a.isFullscreen = false
		}
	})
	
	a.sessionService.AddEventCallback("break_session_start", func() {
		a.audioService.PlayStartSound()
		a.notificationService.ShowBreakSessionStart()
		a.currentScreen = MainScreen
		if a.isFullscreen {
			ebiten.SetFullscreen(false)
			a.isFullscreen = false
		}
	})
	
	a.sessionService.AddEventCallback("work_session_end", func() {
		a.audioService.PlayEndSound()
		a.notificationService.ShowWorkSessionEnd()
		a.currentScreen = FullscreenOverlay
		ebiten.SetFullscreen(true)
		a.isFullscreen = true
		a.setupEndOfWorkButtons()
	})
	
	a.sessionService.AddEventCallback("break_session_end", func() {
		a.audioService.PlayEndSound()
		a.notificationService.ShowBreakSessionEnd()
		a.currentScreen = FullscreenOverlay
		ebiten.SetFullscreen(true)
		a.isFullscreen = true
		a.setupEndOfBreakButtons()
	})
	
	a.sessionService.AddEventCallback("warning", func() {
		a.audioService.PlayWarningSound()
		a.notificationService.ShowWarning()
	})
	
	a.sessionService.AddEventCallback("session_pause", func() {
		a.audioService.PlayBeep(400, 100*time.Millisecond)
		a.notificationService.ShowSessionPaused()
	})
	
	a.sessionService.AddEventCallback("session_resume", func() {
		a.audioService.PlayBeep(600, 100*time.Millisecond)
		a.notificationService.ShowSessionResumed()
	})
}

func (a *App) setupButtons() {
	screenWidth, screenHeight := ebiten.WindowSize()
	
	a.buttons = []Button{
		{
			X: screenWidth/2 - ButtonWidth/2,
			Y: screenHeight/2 - ButtonHeight - ButtonPadding,
			W: ButtonWidth,
			H: ButtonHeight,
			Text: StartWorkButtonText,
			Action: func() {
				a.sessionService.StartWorkSession()
			},
		},
		{
			X: screenWidth/2 - ButtonWidth/2,
			Y: screenHeight/2 + ButtonPadding,
			W: ButtonWidth,
			H: ButtonHeight,
			Text: StartBreakButtonText,
			Action: func() {
				a.sessionService.StartBreakSession()
			},
		},
	}
}

func (a *App) setupEndOfWorkButtons() {
	screenWidth, screenHeight := ebiten.WindowSize()
	
	a.buttons = []Button{
		{
			X: screenWidth/2 - ButtonWidth/2,
			Y: screenHeight/2 - ButtonHeight - ButtonPadding,
			W: ButtonWidth,
			H: ButtonHeight,
			Text: StartBreakButtonText,
			Action: func() {
				a.sessionService.StartBreakSession()
			},
		},
		{
			X: screenWidth/2 - ButtonWidth/2,
			Y: screenHeight/2 + ButtonPadding,
			W: ButtonWidth,
			H: ButtonHeight,
			Text: SkipBreakButtonText,
			Action: func() {
				a.sessionService.StartWorkSession()
			},
		},
	}
}

func (a *App) setupEndOfBreakButtons() {
	screenWidth, screenHeight := ebiten.WindowSize()
	
	a.buttons = []Button{
		{
			X: screenWidth/2 - ButtonWidth/2,
			Y: screenHeight/2,
			W: ButtonWidth,
			H: ButtonHeight,
			Text: StartWorkButtonText,
			Action: func() {
				a.sessionService.StartWorkSession()
			},
		},
	}
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
	a.updateButtons()
	
	return nil
}

func (a *App) updateButtonPositions() {
	screenWidth, screenHeight := ebiten.WindowSize()
	
	// Update button positions based on current screen size
	for i := range a.buttons {
		switch len(a.buttons) {
		case 1: // End of break (single button)
			a.buttons[i].X = screenWidth/2 - ButtonWidth/2
			a.buttons[i].Y = screenHeight/2
		case 2: // Main screen or end of work (two buttons)
			if i == 0 {
				a.buttons[i].X = screenWidth/2 - ButtonWidth/2
				a.buttons[i].Y = screenHeight/2 - ButtonHeight - ButtonPadding
			} else {
				a.buttons[i].X = screenWidth/2 - ButtonWidth/2
				a.buttons[i].Y = screenHeight/2 + ButtonPadding
			}
		}
	}
}

func (a *App) updateButtons() {
	mx, my := ebiten.CursorPosition()
	
	for i := range a.buttons {
		button := &a.buttons[i]
		button.Hovered = mx >= button.X && mx < button.X+button.W &&
			my >= button.Y && my < button.Y+button.H
		
		if button.Hovered && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			button.Action()
		}
	}
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
	timerBg := color.RGBA{R: 0, G: 0, B: 0, A: 100}
	timerX, timerY := screenWidth/2-60, screenHeight/2-110
	a.drawRect(screen, timerX, timerY, 120, 30, timerBg)
	ebitenutil.DebugPrintAt(screen, timerText, screenWidth/2-50, screenHeight/2-100)
	
	// Draw progress bar
	a.drawProgressBar(screen, session.GetProgress(), screenWidth, screenHeight)
	
	if session.IsSessionPaused() {
		ebitenutil.DebugPrintAt(screen, PausedText, screenWidth/2-30, screenHeight/2-50)
		ebitenutil.DebugPrintAt(screen, ResumeInstructionText, screenWidth/2-80, screenHeight/2-20)
	} else {
		ebitenutil.DebugPrintAt(screen, WorkingText, screenWidth/2-30, screenHeight/2-50)
		ebitenutil.DebugPrintAt(screen, PauseInstructionText, screenWidth/2-80, screenHeight/2-20)
	}
}

func (a *App) drawBreakSession(screen *ebiten.Image) {
	session := a.sessionService.GetSession()
	remaining := session.GetTimeRemaining()
	screenWidth, screenHeight := ebiten.WindowSize()
	
	screen.Fill(BreakSessionColor)
	
	timerText := fmt.Sprintf("%02d:%02d", int(remaining.Minutes()), int(remaining.Seconds())%60)
	// Draw timer background for better visibility
	timerBg := color.RGBA{R: 0, G: 0, B: 0, A: 100}
	timerX, timerY := screenWidth/2-60, screenHeight/2-110
	a.drawRect(screen, timerX, timerY, 120, 30, timerBg)
	ebitenutil.DebugPrintAt(screen, timerText, screenWidth/2-50, screenHeight/2-100)
	
	// Draw progress bar
	a.drawProgressBar(screen, session.GetProgress(), screenWidth, screenHeight)
	
	if session.IsSessionPaused() {
		ebitenutil.DebugPrintAt(screen, PausedText, screenWidth/2-30, screenHeight/2-50)
		ebitenutil.DebugPrintAt(screen, ResumeInstructionText, screenWidth/2-80, screenHeight/2-20)
	} else {
		ebitenutil.DebugPrintAt(screen, BreakText, screenWidth/2-40, screenHeight/2-50)
		ebitenutil.DebugPrintAt(screen, PauseInstructionText, screenWidth/2-80, screenHeight/2-20)
	}
}

func (a *App) drawIdleScreen(screen *ebiten.Image) {
	screenWidth, screenHeight := ebiten.WindowSize()
	ebitenutil.DebugPrintAt(screen, IdleScreenMessage, screenWidth/2-100, screenHeight/2-150)
	
	a.drawButtons(screen)
}

func (a *App) drawFullscreenOverlay(screen *ebiten.Image) {
	session := a.sessionService.GetSession()
	screenWidth, screenHeight := ebiten.WindowSize()
	
	// 強制的な赤い背景で注意を引く
	screen.Fill(color.RGBA{R: 180, G: 0, B: 0, A: 255})
	
	var message string
	if session.GetSessionType() == domain.Work {
		message = WorkSessionEndMessage
	} else {
		message = BreakSessionEndMessage
	}
	
	// メッセージを大きく強調表示
	messageX := screenWidth/2 - len(message)*4
	messageY := screenHeight/3
	
	// 背景の強調ボックス
	boxWidth := len(message) * 8 + 40
	boxHeight := 60
	boxX := screenWidth/2 - boxWidth/2
	boxY := messageY - 20
	a.drawRect(screen, boxX, boxY, boxWidth, boxHeight, color.RGBA{R: 255, G: 255, B: 0, A: 200})
	a.drawBorder(screen, boxX, boxY, boxWidth, boxHeight, color.RGBA{R: 255, G: 255, B: 255, A: 255}, 3)
	
	ebitenutil.DebugPrintAt(screen, message, messageX, messageY)
	
	// 警告メッセージを追加
	warningMsg := "YOU CANNOT CONTINUE UNTIL YOU CHOOSE!"
	warningX := screenWidth/2 - len(warningMsg)*3
	warningY := screenHeight/2 - 50
	ebitenutil.DebugPrintAt(screen, warningMsg, warningX, warningY)
	
	a.drawButtons(screen)
}

func (a *App) drawButtons(screen *ebiten.Image) {
	for _, button := range a.buttons {
		// Draw button shadow
		shadowColor := color.RGBA{R: 0, G: 0, B: 0, A: 50}
		a.drawRect(screen, button.X+2, button.Y+2, button.W, button.H, shadowColor)
		
		// Draw button background
		buttonColor := ButtonColor
		if button.Hovered {
			buttonColor = ButtonHoverColor
		}
		a.drawRect(screen, button.X, button.Y, button.W, button.H, buttonColor)
		
		// Draw button border
		borderColor := color.RGBA{R: 200, G: 200, B: 200, A: 255}
		a.drawBorder(screen, button.X, button.Y, button.W, button.H, borderColor, 2)
		
		// Draw button text (centered)
		textX := button.X + button.W/2 - len(button.Text)*3
		textY := button.Y + button.H/2 - 6
		ebitenutil.DebugPrintAt(screen, button.Text, textX, textY)
	}
}

func (a *App) drawRect(screen *ebiten.Image, x, y, w, h int, c color.Color) {
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			screen.Set(x+i, y+j, c)
		}
	}
}

func (a *App) drawBorder(screen *ebiten.Image, x, y, w, h int, c color.Color, thickness int) {
	// Top border
	a.drawRect(screen, x, y, w, thickness, c)
	// Bottom border
	a.drawRect(screen, x, y+h-thickness, w, thickness, c)
	// Left border
	a.drawRect(screen, x, y, thickness, h, c)
	// Right border
	a.drawRect(screen, x+w-thickness, y, thickness, h, c)
}

func (a *App) drawProgressBar(screen *ebiten.Image, progress float64, screenWidth, screenHeight int) {
	barWidth := 300
	barHeight := 10
	barX := screenWidth/2 - barWidth/2
	barY := screenHeight/2 + 40
	
	// Draw progress bar background
	bgColor := color.RGBA{R: 60, G: 60, B: 60, A: 255}
	a.drawRect(screen, barX, barY, barWidth, barHeight, bgColor)
	
	// Draw progress bar fill
	progressWidth := int(float64(barWidth) * progress)
	progressColor := color.RGBA{R: 100, G: 200, B: 100, A: 255}
	if progressWidth > 0 {
		a.drawRect(screen, barX, barY, progressWidth, barHeight, progressColor)
	}
	
	// Draw progress bar border
	borderColor := color.RGBA{R: 180, G: 180, B: 180, A: 255}
	a.drawBorder(screen, barX, barY, barWidth, barHeight, borderColor, 1)
}

func (a *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// Allow dynamic resizing but set minimum size
	minWidth, minHeight := 600, 400
	
	if outsideWidth < minWidth {
		outsideWidth = minWidth
	}
	if outsideHeight < minHeight {
		outsideHeight = minHeight
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