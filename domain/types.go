package domain

import "time"

const (
	WorkSessionDuration    = 25 * time.Minute
	BreakSessionDuration   = 5 * time.Minute
	WarningInterval        = 5 * time.Minute
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