package application

import (
	"karedoro/domain"
)

// Services provides a container for all application services with dependency injection.
type Services struct {
	Session      *SessionService
	Audio        domain.AudioPlayer
	Notification domain.NotificationSender
	Config       *ConfigService
}

// NewServices creates a new Services container with all dependencies wired up.
func NewServices() *Services {
	audioService := NewAudioService()
	notificationService := NewNotificationService()
	sessionService := NewSessionService()
	configService := NewConfigService()
	
	return &Services{
		Session:      sessionService,
		Audio:        audioService,
		Notification: notificationService,
		Config:       configService,
	}
}

// NewServicesWithDependencies creates a new Services container with injected dependencies.
// This is useful for testing or when you need custom implementations of interfaces.
func NewServicesWithDependencies(
	audio domain.AudioPlayer,
	notification domain.NotificationSender,
) *Services {
	sessionService := NewSessionService()
	configService := NewConfigService()
	
	return &Services{
		Session:      sessionService,
		Audio:        audio,
		Notification: notification,
		Config:       configService,
	}
}