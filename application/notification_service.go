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
		"WORK SESSION STARTED! Focus for 25 minutes - NO DISTRACTIONS!",
		"",
	)
}

func (n *NotificationService) ShowBreakSessionStart() error {
	if !n.enabled {
		return nil
	}
	
	return beeep.Notify(
		n.appName,
		"BREAK SESSION STARTED! Relax for 5 minutes - You earned it!",
		"",
	)
}

func (n *NotificationService) ShowWorkSessionEnd() error {
	if !n.enabled {
		return nil
	}
	
	return beeep.Notify(
		n.appName,
		"POMODORO COMPLETE! You MUST take a break now - No skipping!",
		"",
	)
}

func (n *NotificationService) ShowBreakSessionEnd() error {
	if !n.enabled {
		return nil
	}
	
	return beeep.Notify(
		n.appName,
		"BREAK OVER! Time to get back to work - Start your session NOW!",
		"",
	)
}

func (n *NotificationService) ShowWarning() error {
	if !n.enabled {
		return nil
	}
	
	return beeep.Notify(
		n.appName,
		"WARNING! You haven't started your next session! FOLLOW THE POMODORO TECHNIQUE!",
		"",
	)
}

func (n *NotificationService) ShowSessionPaused() error {
	if !n.enabled {
		return nil
	}
	
	return beeep.Notify(
		n.appName,
		"Session paused",
		"",
	)
}

func (n *NotificationService) ShowSessionResumed() error {
	if !n.enabled {
		return nil
	}
	
	return beeep.Notify(
		n.appName,
		"Session resumed",
		"",
	)
}

func (n *NotificationService) ShowCustomMessage(title, message string) error {
	if !n.enabled {
		return nil
	}
	
	return beeep.Notify(title, message, "")
}