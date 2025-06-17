package presentation

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"karedoro/application"
	"karedoro/domain"
)

// AppCoordinator coordinates between different components of the application.
type AppCoordinator struct {
	sessionService *application.SessionService
	configService  *application.ConfigService
	eventHandler   *EventHandler
	uiManager      *UIManager
	inputHandler   *InputHandler
}

func NewAppCoordinator(sessionService *application.SessionService, configService *application.ConfigService, eventHandler *EventHandler) *AppCoordinator {
	coordinator := &AppCoordinator{
		sessionService: sessionService,
		configService:  configService,
		eventHandler:   eventHandler,
		uiManager:      NewUIManager(),
		inputHandler:   NewInputHandler(sessionService),
	}
	
	coordinator.setupEventCallbacks()
	
	return coordinator
}

func (ac *AppCoordinator) setupEventCallbacks() {
	ac.eventHandler.SetupCallbacks(
		ac.sessionService,
		func() {
			ac.uiManager.SetCurrentScreen(FullscreenOverlay)
			ac.uiManager.SetFullscreen(true)
			screenWidth, screenHeight := ebiten.WindowSize()
			ac.uiManager.SetupEndOfWorkButtons(screenWidth, screenHeight, ac.sessionService)
		},
		func() {
			ac.uiManager.SetCurrentScreen(FullscreenOverlay)
			ac.uiManager.SetFullscreen(true)
			screenWidth, screenHeight := ebiten.WindowSize()
			ac.uiManager.SetupEndOfBreakButtons(screenWidth, screenHeight, ac.sessionService)
		},
	)
	
	// Handle session start events that need screen changes
	ac.sessionService.AddEventCallback(domain.EventWorkSessionStart, func() {
		ac.uiManager.SetCurrentScreen(MainScreen)
		if ac.uiManager.IsFullscreen() {
			ebiten.SetFullscreen(false)
			ac.uiManager.SetFullscreen(false)
		}
	})
	
	ac.sessionService.AddEventCallback(domain.EventBreakSessionStart, func() {
		ac.uiManager.SetCurrentScreen(MainScreen)
		if ac.uiManager.IsFullscreen() {
			ebiten.SetFullscreen(false)
			ac.uiManager.SetFullscreen(false)
		}
	})
}

func (ac *AppCoordinator) Initialize() {
	// Setup initial buttons
	screenWidth, screenHeight := ebiten.WindowSize()
	ac.uiManager.SetupMainButtons(screenWidth, screenHeight, ac.sessionService)
}

func (ac *AppCoordinator) Update() error {
	ac.sessionService.Update()
	ac.inputHandler.HandleInput()
	
	// Update button positions and handle interactions
	screenWidth, screenHeight := ebiten.WindowSize()
	ac.uiManager.UpdateButtonPositions(screenWidth, screenHeight)
	ac.uiManager.UpdateButtons()
	
	return nil
}

func (ac *AppCoordinator) Draw(screen *ebiten.Image) {
	ac.uiManager.Draw(screen, ac.sessionService.GetSession())
}

func (ac *AppCoordinator) RunSetup(audioService *application.AudioService) error {
	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowTitle(WindowTitle)
	ebiten.SetWindowClosingHandled(true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	
	if !audioService.WaitForReady(5 * time.Second) {
		fmt.Println("Warning: Audio system initialization timeout")
	}
	
	return nil
}

func (ac *AppCoordinator) RunSetupWithServices() error {
	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowTitle(WindowTitle)
	ebiten.SetWindowClosingHandled(true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	
	// Wait for audio service to be ready if it's an AudioService
	if audioService, ok := ac.eventHandler.audioService.(*application.AudioService); ok {
		if !audioService.WaitForReady(5 * time.Second) {
			fmt.Println("Warning: Audio system initialization timeout")
		}
	}
	
	return nil
}