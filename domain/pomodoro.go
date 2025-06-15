package domain

import (
	"context"
	"sync"
	"time"
)

// PomodoroService defines the interface for pomodoro operations
type PomodoroService interface {
	StartWorkSession() error
	StartBreakSession() error
	PauseSession() error
	ResumeSession() error
	GetCurrentSession() *Session
	GetCurrentState() SessionState
	GetRemainingTime() time.Duration
	IsSessionExpired() bool
}

// EventHandler defines the interface for handling pomodoro events
type EventHandler interface {
	OnSessionStart(session *Session)
	OnSessionEnd(session *Session)
	OnSessionPause(session *Session)
	OnSessionResume(session *Session)
	OnTimerTick(session *Session, remainingTime time.Duration)
	OnWarning(idleDuration time.Duration)
}

// PomodoroManager manages the pomodoro sessions and state
type PomodoroManager struct {
	config       *SessionConfig
	currentSession *Session
	state        SessionState
	eventHandler EventHandler
	ctx          context.Context
	cancel       context.CancelFunc
	mu           sync.RWMutex
	lastStateChange time.Time
}

// NewPomodoroManager creates a new PomodoroManager
func NewPomodoroManager(config *SessionConfig, handler EventHandler) *PomodoroManager {
	ctx, cancel := context.WithCancel(context.Background())
	
	return &PomodoroManager{
		config:       config,
		state:        Idle,
		eventHandler: handler,
		ctx:          ctx,
		cancel:       cancel,
		lastStateChange: time.Now(),
	}
}

// StartWorkSession starts a new work session
func (pm *PomodoroManager) StartWorkSession() error {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	
	pm.currentSession = NewWorkSession(pm.config)
	pm.state = WorkSession
	pm.lastStateChange = time.Now()
	
	if pm.eventHandler != nil {
		pm.eventHandler.OnSessionStart(pm.currentSession)
	}
	
	go pm.sessionTimer()
	
	return nil
}

// StartBreakSession starts a new break session
func (pm *PomodoroManager) StartBreakSession() error {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	
	pm.currentSession = NewBreakSession(pm.config)
	pm.state = BreakSession
	pm.lastStateChange = time.Now()
	
	if pm.eventHandler != nil {
		pm.eventHandler.OnSessionStart(pm.currentSession)
	}
	
	go pm.sessionTimer()
	
	return nil
}

// PauseSession pauses the current session
func (pm *PomodoroManager) PauseSession() error {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	
	if pm.currentSession != nil && !pm.currentSession.IsPaused {
		pm.currentSession.Pause()
		
		if pm.eventHandler != nil {
			pm.eventHandler.OnSessionPause(pm.currentSession)
		}
	}
	
	return nil
}

// ResumeSession resumes the current session
func (pm *PomodoroManager) ResumeSession() error {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	
	if pm.currentSession != nil && pm.currentSession.IsPaused {
		pm.currentSession.Resume()
		
		if pm.eventHandler != nil {
			pm.eventHandler.OnSessionResume(pm.currentSession)
		}
	}
	
	return nil
}

// GetCurrentSession returns the current session
func (pm *PomodoroManager) GetCurrentSession() *Session {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	
	return pm.currentSession
}

// GetCurrentState returns the current state
func (pm *PomodoroManager) GetCurrentState() SessionState {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	
	return pm.state
}

// GetRemainingTime returns the remaining time of the current session
func (pm *PomodoroManager) GetRemainingTime() time.Duration {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	
	if pm.currentSession != nil {
		return pm.currentSession.GetRemainingTime()
	}
	
	return 0
}

// IsSessionExpired returns true if the current session has expired
func (pm *PomodoroManager) IsSessionExpired() bool {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	
	if pm.currentSession != nil {
		return pm.currentSession.IsExpired()
	}
	
	return false
}

// GetIdleDuration returns how long the app has been idle
func (pm *PomodoroManager) GetIdleDuration() time.Duration {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	
	if pm.state == Idle {
		return time.Since(pm.lastStateChange)
	}
	
	return 0
}

// sessionTimer runs the session timer in a separate goroutine
func (pm *PomodoroManager) sessionTimer() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	
	for {
		select {
		case <-pm.ctx.Done():
			return
		case <-ticker.C:
			pm.mu.RLock()
			session := pm.currentSession
			pm.mu.RUnlock()
			
			if session == nil {
				return
			}
			
			remainingTime := session.GetRemainingTime()
			
			if pm.eventHandler != nil && !session.IsPaused {
				pm.eventHandler.OnTimerTick(session, remainingTime)
			}
			
			if session.IsExpired() {
				pm.endSession()
				return
			}
		}
	}
}

// endSession handles session completion
func (pm *PomodoroManager) endSession() {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	
	if pm.currentSession != nil {
		if pm.eventHandler != nil {
			pm.eventHandler.OnSessionEnd(pm.currentSession)
		}
		
		pm.currentSession = nil
		pm.state = Idle
		pm.lastStateChange = time.Now()
		
		go pm.warningTimer()
	}
}

// warningTimer handles idle warnings
func (pm *PomodoroManager) warningTimer() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	
	for {
		select {
		case <-pm.ctx.Done():
			return
		case <-ticker.C:
			pm.mu.RLock()
			state := pm.state
			idleDuration := pm.GetIdleDuration()
			pm.mu.RUnlock()
			
			if state != Idle {
				return
			}
			
			if pm.eventHandler != nil {
				pm.eventHandler.OnWarning(idleDuration)
			}
		}
	}
}

// Shutdown shuts down the PomodoroManager
func (pm *PomodoroManager) Shutdown() {
	pm.cancel()
}