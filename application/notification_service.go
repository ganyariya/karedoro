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
		"作業セッション開始！集中して取り組みましょう。",
		"",
	)
}

func (n *NotificationService) ShowBreakSessionStart() error {
	if !n.enabled {
		return nil
	}
	
	return beeep.Notify(
		n.appName,
		"休憩セッション開始！リラックスしてください。",
		"",
	)
}

func (n *NotificationService) ShowWorkSessionEnd() error {
	if !n.enabled {
		return nil
	}
	
	return beeep.Notify(
		n.appName,
		"ポモドーロ完了！休憩しましょう！",
		"",
	)
}

func (n *NotificationService) ShowBreakSessionEnd() error {
	if !n.enabled {
		return nil
	}
	
	return beeep.Notify(
		n.appName,
		"休憩終了！作業に戻りましょう！",
		"",
	)
}

func (n *NotificationService) ShowWarning() error {
	if !n.enabled {
		return nil
	}
	
	return beeep.Notify(
		n.appName,
		"まだ次のセッションを開始していません！",
		"",
	)
}

func (n *NotificationService) ShowSessionPaused() error {
	if !n.enabled {
		return nil
	}
	
	return beeep.Notify(
		n.appName,
		"セッションが一時停止されました。",
		"",
	)
}

func (n *NotificationService) ShowSessionResumed() error {
	if !n.enabled {
		return nil
	}
	
	return beeep.Notify(
		n.appName,
		"セッションが再開されました。",
		"",
	)
}

func (n *NotificationService) ShowCustomMessage(title, message string) error {
	if !n.enabled {
		return nil
	}
	
	return beeep.Notify(title, message, "")
}