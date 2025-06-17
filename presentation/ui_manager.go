package presentation

import (
	"github.com/hajimehoshi/ebiten/v2"
	"karedoro/application"
	"karedoro/domain"
)

// UIManager manages the overall UI state and coordinates screen rendering.
type UIManager struct {
	currentScreen   Screen
	isFullscreen    bool
	buttonManager   *ButtonManager
	screenRenderer  *ScreenRenderer
}

func NewUIManager() *UIManager {
	return &UIManager{
		currentScreen:  MainScreen,
		isFullscreen:   false,
		buttonManager:  NewButtonManager(),
		screenRenderer: NewScreenRenderer(),
	}
}

func (ui *UIManager) SetCurrentScreen(screen Screen) {
	ui.currentScreen = screen
}

func (ui *UIManager) GetCurrentScreen() Screen {
	return ui.currentScreen
}

func (ui *UIManager) SetFullscreen(fullscreen bool) {
	ui.isFullscreen = fullscreen
}

func (ui *UIManager) IsFullscreen() bool {
	return ui.isFullscreen
}

func (ui *UIManager) GetButtonManager() *ButtonManager {
	return ui.buttonManager
}

func (ui *UIManager) GetScreenRenderer() *ScreenRenderer {
	return ui.screenRenderer
}

func (ui *UIManager) UpdateButtonPositions(screenWidth, screenHeight int) {
	ui.buttonManager.UpdateButtonPositions(screenWidth, screenHeight)
}

func (ui *UIManager) UpdateButtons() {
	ui.buttonManager.UpdateButtons()
}

func (ui *UIManager) Draw(screen *ebiten.Image, session *domain.Session) {
	screen.Fill(BackgroundColor)
	
	switch ui.currentScreen {
	case MainScreen:
		ui.screenRenderer.DrawMainScreen(screen, session, ui.buttonManager)
	case FullscreenOverlay:
		ui.screenRenderer.DrawFullscreenOverlay(screen, session, ui.buttonManager)
	}
}

func (ui *UIManager) SetupMainButtons(screenWidth, screenHeight int, sessionService *application.SessionService) {
	ui.buttonManager.SetupMainButtons(screenWidth, screenHeight, sessionService)
}

func (ui *UIManager) SetupEndOfWorkButtons(screenWidth, screenHeight int, sessionService *application.SessionService) {
	ui.buttonManager.SetupEndOfWorkButtons(screenWidth, screenHeight, sessionService)
}

func (ui *UIManager) SetupEndOfBreakButtons(screenWidth, screenHeight int, sessionService *application.SessionService) {
	ui.buttonManager.SetupEndOfBreakButtons(screenWidth, screenHeight, sessionService)
}