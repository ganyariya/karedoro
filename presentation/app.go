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
	a.buttons = []Button{
		{
			X: WindowWidth/2 - ButtonWidth/2,
			Y: WindowHeight/2 - ButtonHeight - ButtonPadding,
			W: ButtonWidth,
			H: ButtonHeight,
			Text: StartWorkButtonText,
			Action: func() {
				a.sessionService.StartWorkSession()
			},
		},
		{
			X: WindowWidth/2 - ButtonWidth/2,
			Y: WindowHeight/2 + ButtonPadding,
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
	a.buttons = []Button{
		{
			X: WindowWidth/2 - ButtonWidth/2,
			Y: WindowHeight/2 - ButtonHeight - ButtonPadding,
			W: ButtonWidth,
			H: ButtonHeight,
			Text: StartBreakButtonText,
			Action: func() {
				a.sessionService.StartBreakSession()
			},
		},
		{
			X: WindowWidth/2 - ButtonWidth/2,
			Y: WindowHeight/2 + ButtonPadding,
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
	a.buttons = []Button{
		{
			X: WindowWidth/2 - ButtonWidth/2,
			Y: WindowHeight/2,
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
	
	a.updateButtons()
	
	return nil
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
	
	screen.Fill(WorkSessionColor)
	
	timerText := fmt.Sprintf("%02d:%02d", int(remaining.Minutes()), int(remaining.Seconds())%60)
	ebitenutil.DebugPrintAt(screen, timerText, WindowWidth/2-50, WindowHeight/2-100)
	
	if session.IsSessionPaused() {
		ebitenutil.DebugPrintAt(screen, PausedText, WindowWidth/2-40, WindowHeight/2-50)
		ebitenutil.DebugPrintAt(screen, "スペースキーで再開", WindowWidth/2-60, WindowHeight/2-20)
	} else {
		ebitenutil.DebugPrintAt(screen, "作業中", WindowWidth/2-30, WindowHeight/2-50)
		ebitenutil.DebugPrintAt(screen, "スペースキーで一時停止", WindowWidth/2-80, WindowHeight/2-20)
	}
}

func (a *App) drawBreakSession(screen *ebiten.Image) {
	session := a.sessionService.GetSession()
	remaining := session.GetTimeRemaining()
	
	screen.Fill(BreakSessionColor)
	
	timerText := fmt.Sprintf("%02d:%02d", int(remaining.Minutes()), int(remaining.Seconds())%60)
	ebitenutil.DebugPrintAt(screen, timerText, WindowWidth/2-50, WindowHeight/2-100)
	
	if session.IsSessionPaused() {
		ebitenutil.DebugPrintAt(screen, PausedText, WindowWidth/2-40, WindowHeight/2-50)
		ebitenutil.DebugPrintAt(screen, "スペースキーで再開", WindowWidth/2-60, WindowHeight/2-20)
	} else {
		ebitenutil.DebugPrintAt(screen, "休憩中", WindowWidth/2-30, WindowHeight/2-50)
		ebitenutil.DebugPrintAt(screen, "スペースキーで一時停止", WindowWidth/2-80, WindowHeight/2-20)
	}
}

func (a *App) drawIdleScreen(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "次のセッションを選択してください", WindowWidth/2-120, WindowHeight/2-150)
	
	a.drawButtons(screen)
}

func (a *App) drawFullscreenOverlay(screen *ebiten.Image) {
	session := a.sessionService.GetSession()
	
	var message string
	if session.GetSessionType() == domain.Work {
		message = WorkSessionEndMessage
	} else {
		message = BreakSessionEndMessage
	}
	
	ebitenutil.DebugPrintAt(screen, message, WindowWidth/2-len(message)*3, WindowHeight/3)
	
	a.drawButtons(screen)
}

func (a *App) drawButtons(screen *ebiten.Image) {
	for _, button := range a.buttons {
		buttonColor := ButtonColor
		if button.Hovered {
			buttonColor = ButtonHoverColor
		}
		
		a.drawRect(screen, button.X, button.Y, button.W, button.H, buttonColor)
		
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

func (a *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WindowWidth, WindowHeight
}

func (a *App) Run() error {
	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowTitle(WindowTitle)
	ebiten.SetWindowClosingHandled(true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	
	if !a.audioService.WaitForReady(5 * time.Second) {
		fmt.Println("Warning: Audio system initialization timeout")
	}
	
	return ebiten.RunGame(a)
}

func (a *App) IsDrawingSkipped() bool {
	return false
}