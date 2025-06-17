package domain

import (
	"errors"
	"fmt"
)

// SessionError represents errors related to session operations.
type SessionError struct {
	Op  string
	Err error
}

func (e *SessionError) Error() string {
	return fmt.Sprintf("session %s: %v", e.Op, e.Err)
}

func (e *SessionError) Unwrap() error {
	return e.Err
}

// NewSessionError creates a new SessionError with the given operation and underlying error.
func NewSessionError(op string, err error) *SessionError {
	return &SessionError{
		Op:  op,
		Err: err,
	}
}

// TimerError represents errors related to timer operations.
type TimerError struct {
	Op  string
	Err error
}

func (e *TimerError) Error() string {
	return fmt.Sprintf("timer %s: %v", e.Op, e.Err)
}

func (e *TimerError) Unwrap() error {
	return e.Err
}

// NewTimerError creates a new TimerError with the given operation and underlying error.
func NewTimerError(op string, err error) *TimerError {
	return &TimerError{
		Op:  op,
		Err: err,
	}
}

// Common domain errors.
var (
	ErrInvalidState      = errors.New("invalid session state")
	ErrTimerNotRunning   = errors.New("timer is not running")
	ErrTimerAlreadyStarted = errors.New("timer is already started")
	ErrInvalidDuration   = errors.New("invalid duration")
	ErrSessionNotFound   = errors.New("session not found")
	ErrConfigNotFound    = errors.New("configuration not found")
	ErrInvalidConfig     = errors.New("invalid configuration")
)

// Audio service errors.
var (
	ErrAudioNotReady     = errors.New("audio service not ready")
	ErrAudioPlayback     = errors.New("audio playback failed")
	ErrAudioInitialization = errors.New("audio initialization failed")
)

// Notification service errors.
var (
	ErrNotificationFailed = errors.New("notification failed")
)

// Configuration errors.
var (
	ErrConfigLoad = errors.New("failed to load configuration")
	ErrConfigSave = errors.New("failed to save configuration")
)