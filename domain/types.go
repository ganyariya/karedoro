package domain

import "time"

const (
	WorkSessionDuration    = 25 * time.Minute
	BreakSessionDuration   = 5 * time.Minute
	WarningInterval        = 5 * time.Minute
)

// Event name constants
const (
	EventWorkSessionStart  = "work_session_start"
	EventBreakSessionStart = "break_session_start"
	EventWorkSessionEnd    = "work_session_end"
	EventBreakSessionEnd   = "break_session_end"
	EventWarning           = "warning"
	EventSessionPause      = "session_pause"
	EventSessionResume     = "session_resume"
)

type SessionState int

const (
	WorkSession SessionState = iota
	BreakSession
	Idle
)

func (s SessionState) String() string {
	switch s {
	case WorkSession:
		return "WorkSession"
	case BreakSession:
		return "BreakSession"
	case Idle:
		return "Idle"
	default:
		return "Unknown"
	}
}

type SessionType int

const (
	Work SessionType = iota
	Break
)

func (t SessionType) String() string {
	switch t {
	case Work:
		return "Work"
	case Break:
		return "Break"
	default:
		return "Unknown"
	}
}

func (t SessionType) Duration() time.Duration {
	switch t {
	case Work:
		return WorkSessionDuration
	case Break:
		return BreakSessionDuration
	default:
		return 0
	}
}