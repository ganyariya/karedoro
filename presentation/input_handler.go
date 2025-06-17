package presentation

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"karedoro/application"
	"karedoro/domain"
)

// InputHandler manages user input processing.
type InputHandler struct {
	sessionService *application.SessionService
}

func NewInputHandler(sessionService *application.SessionService) *InputHandler {
	return &InputHandler{
		sessionService: sessionService,
	}
}

func (ih *InputHandler) HandleInput() {
	session := ih.sessionService.GetSession()
	
	switch session.GetState() {
	case domain.WorkSession, domain.BreakSession:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			if session.IsSessionPaused() {
				ih.sessionService.ResumeSession()
			} else {
				ih.sessionService.PauseSession()
			}
		}
	}
}