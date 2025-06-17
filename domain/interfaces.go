package domain

import "time"

// AudioPlayer provides audio playback functionality for the Pomodoro timer.
type AudioPlayer interface {
	PlayStartSound() error
	PlayEndSound() error
	PlayWarningSound() error
	PlayBeep(frequency float64, duration time.Duration) error
	IsReady() bool
}

// NotificationSender provides system notification functionality.
type NotificationSender interface {
	ShowWorkSessionStart() error
	ShowBreakSessionStart() error
	ShowWorkSessionEnd() error
	ShowBreakSessionEnd() error
	ShowWarning() error
	ShowSessionPaused() error
	ShowSessionResumed() error
}

// ConfigRepository handles configuration persistence.
type ConfigRepository interface {
	Load() (*Config, error)
	Save(config *Config) error
}

// Config represents the application configuration.
type Config struct {
	WorkDuration  time.Duration `json:"work_duration"`
	BreakDuration time.Duration `json:"break_duration"`
}

// SessionRepository handles session state persistence.
type SessionRepository interface {
	Save(session *Session) error
	Load() (*Session, error)
}