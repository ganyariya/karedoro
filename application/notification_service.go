package application

import (
	"github.com/gen2brain/beeep"
)

type NotificationService struct {
	appName string
	enabled bool
}

func NewNotificationService() *NotificationService {
	return &NotificationService{
		appName: "karedoro",
		enabled: true,
	}
}

func (n *NotificationService) SetEnabled(enabled bool) {
	n.enabled = enabled
}

func (n *NotificationService) IsEnabled() bool {
	return n.enabled
}

func (n *NotificationService) ShowWorkSessionStart() error {
	if !n.enabled {
		return nil
	}
	
	return beeep.Notify(
		n.appName,
		"Work Session Started! Time to focus.",
		"",
	)
}

func (n *NotificationService) ShowBreakSessionStart() error {
	if !n.enabled {
		return nil
	}
	
	return beeep.Notify(
		n.appName,
		"Break Session Started! Time to relax.",
		"",
	)
}

func (n *NotificationService) ShowWorkSessionEnd() error {
	if !n.enabled {
		return nil
	}
	
	return beeep.Notify(
		n.appName,
		"Pomodoro Complete! Take a break!",
		"",
	)
}

func (n *NotificationService) ShowBreakSessionEnd() error {
	if !n.enabled {
		return nil
	}
	
	return beeep.Notify(
		n.appName,
		"Break Over! Back to work!",
		"",
	)
}

func (n *NotificationService) ShowWarning() error {
	if !n.enabled {
		return nil
	}
	
	return beeep.Notify(
		n.appName,
		"Haven't started next session yet!",
		"",
	)
}

func (n *NotificationService) ShowSessionPaused() error {
	if !n.enabled {
		return nil
	}
	
	return beeep.Notify(
		n.appName,
		"Session paused.",
		"",
	)
}

func (n *NotificationService) ShowSessionResumed() error {
	if !n.enabled {
		return nil
	}
	
	return beeep.Notify(
		n.appName,
		"Session resumed.",
		"",
	)
}

func (n *NotificationService) ShowCustomMessage(title, message string) error {
	if !n.enabled {
		return nil
	}
	
	return beeep.Notify(title, message, "")
}