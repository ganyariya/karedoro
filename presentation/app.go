package presentation

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
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
	screenRenderer  *ScreenRenderer
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
		screenRenderer:      NewScreenRenderer(),
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
		a.screenRenderer.DrawMainScreen(screen, a.sessionService.GetSession(), a.buttonManager)
	case FullscreenOverlay:
		a.screenRenderer.DrawFullscreenOverlay(screen, a.sessionService.GetSession(), a.buttonManager)
	}
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