package presentation

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"karedoro/application"
	"karedoro/domain"
)

type EventHandler struct {
	audioService        *application.AudioService
	notificationService *application.NotificationService
}

func NewEventHandler(audioService *application.AudioService, notificationService *application.NotificationService) *EventHandler {
	return &EventHandler{
		audioService:        audioService,
		notificationService: notificationService,
	}
}

func (eh *EventHandler) SetupCallbacks(sessionService *application.SessionService, onWorkSessionEnd, onBreakSessionEnd func()) {
	sessionService.AddEventCallback(domain.EventWorkSessionStart, func() {
		eh.audioService.PlayStartSound()
		eh.notificationService.ShowWorkSessionStart()
	})
	
	sessionService.AddEventCallback(domain.EventBreakSessionStart, func() {
		eh.audioService.PlayStartSound()
		eh.notificationService.ShowBreakSessionStart()
	})
	
	sessionService.AddEventCallback(domain.EventWorkSessionEnd, func() {
		eh.audioService.PlayEndSound()
		eh.notificationService.ShowWorkSessionEnd()
		ebiten.SetFullscreen(true)
		onWorkSessionEnd()
	})
	
	sessionService.AddEventCallback(domain.EventBreakSessionEnd, func() {
		eh.audioService.PlayEndSound()
		eh.notificationService.ShowBreakSessionEnd()
		ebiten.SetFullscreen(true)
		onBreakSessionEnd()
	})
	
	sessionService.AddEventCallback(domain.EventWarning, func() {
		eh.audioService.PlayWarningSound()
		eh.notificationService.ShowWarning()
	})
	
	sessionService.AddEventCallback(domain.EventSessionPause, func() {
		eh.audioService.PlayBeep(400, 100*time.Millisecond)
		eh.notificationService.ShowSessionPaused()
	})
	
	sessionService.AddEventCallback(domain.EventSessionResume, func() {
		eh.audioService.PlayBeep(600, 100*time.Millisecond)
		eh.notificationService.ShowSessionResumed()
	})
}