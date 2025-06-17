package presentation

import (
	"github.com/hajimehoshi/ebiten/v2"

	"karedoro/application"
)

type App struct {
	coordinator *AppCoordinator
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

func NewApp() (*App, *application.AudioService) {
	audioService := application.NewAudioService()
	notificationService := application.NewNotificationService()
	sessionService := application.NewSessionService()
	configService := application.NewConfigService()
	eventHandler := NewEventHandler(audioService, notificationService)
	
	coordinator := NewAppCoordinator(sessionService, configService, eventHandler)
	coordinator.Initialize()
	
	app := &App{
		coordinator: coordinator,
	}
	
	return app, audioService
}

// NewAppWithServices creates a new App with dependency injection.
func NewAppWithServices(services *application.Services) *App {
	eventHandler := NewEventHandler(services.Audio, services.Notification)
	coordinator := NewAppCoordinator(services.Session, services.Config, eventHandler)
	coordinator.Initialize()
	
	return &App{
		coordinator: coordinator,
	}
}




func (a *App) Update() error {
	return a.coordinator.Update()
}


func (a *App) Draw(screen *ebiten.Image) {
	a.coordinator.Draw(screen)
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

func (a *App) RunWithAudioService(audioService *application.AudioService) error {
	if err := a.coordinator.RunSetup(audioService); err != nil {
		return err
	}
	
	return ebiten.RunGame(a)
}

// Run runs the application with dependency injection (no audio service parameter needed).
func (a *App) Run() error {
	if err := a.coordinator.RunSetupWithServices(); err != nil {
		return err
	}
	
	return ebiten.RunGame(a)
}

func (a *App) IsDrawingSkipped() bool {
	return false
}