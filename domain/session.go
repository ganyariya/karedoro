package domain

import (
	"time"
)

// SessionState represents the current state of the pomodoro application
type SessionState int

const (
	Idle SessionState = iota
	WorkSession
	BreakSession
)

func (s SessionState) String() string {
	switch s {
	case Idle:
		return "Idle"
	case WorkSession:
		return "WorkSession"
	case BreakSession:
		return "BreakSession"
	default:
		return "Unknown"
	}
}

// SessionConfig holds the configuration for pomodoro sessions
type SessionConfig struct {
	WorkDuration  time.Duration
	BreakDuration time.Duration
}

// NewDefaultConfig creates a new SessionConfig with default values
func NewDefaultConfig() *SessionConfig {
	return &SessionConfig{
		WorkDuration:  25 * time.Minute,
		BreakDuration: 5 * time.Minute,
	}
}

// Session represents a pomodoro session
type Session struct {
	State     SessionState
	StartTime time.Time
	Duration  time.Duration
	IsPaused  bool
	PausedAt  time.Time
	PauseDuration time.Duration
}

// NewWorkSession creates a new work session
func NewWorkSession(config *SessionConfig) *Session {
	return &Session{
		State:     WorkSession,
		StartTime: time.Now(),
		Duration:  config.WorkDuration,
		IsPaused:  false,
	}
}

// NewBreakSession creates a new break session
func NewBreakSession(config *SessionConfig) *Session {
	return &Session{
		State:     BreakSession,
		StartTime: time.Now(),
		Duration:  config.BreakDuration,
		IsPaused:  false,
	}
}

// GetRemainingTime returns the remaining time in the session
func (s *Session) GetRemainingTime() time.Duration {
	if s.IsPaused {
		elapsed := s.PausedAt.Sub(s.StartTime) - s.PauseDuration
		return s.Duration - elapsed
	}
	
	elapsed := time.Since(s.StartTime) - s.PauseDuration
	remaining := s.Duration - elapsed
	
	if remaining < 0 {
		return 0
	}
	
	return remaining
}

// IsExpired returns true if the session has expired
func (s *Session) IsExpired() bool {
	return s.GetRemainingTime() == 0
}

// Pause pauses the session
func (s *Session) Pause() {
	if !s.IsPaused {
		s.IsPaused = true
		s.PausedAt = time.Now()
	}
}

// Resume resumes the session
func (s *Session) Resume() {
	if s.IsPaused {
		s.PauseDuration += time.Since(s.PausedAt)
		s.IsPaused = false
		s.PausedAt = time.Time{}
	}
}