package domain

import (
	"testing"
	"time"
)

// MockEventHandler for testing
type MockEventHandler struct {
	SessionStartCalled  bool
	SessionEndCalled    bool
	SessionPauseCalled  bool
	SessionResumeCalled bool
	TimerTickCalled     bool
	WarningCalled       bool
	LastSession         *Session
	LastRemainingTime   time.Duration
	LastIdleDuration    time.Duration
}

func (m *MockEventHandler) OnSessionStart(session *Session) {
	m.SessionStartCalled = true
	m.LastSession = session
}

func (m *MockEventHandler) OnSessionEnd(session *Session) {
	m.SessionEndCalled = true
	m.LastSession = session
}

func (m *MockEventHandler) OnSessionPause(session *Session) {
	m.SessionPauseCalled = true
	m.LastSession = session
}

func (m *MockEventHandler) OnSessionResume(session *Session) {
	m.SessionResumeCalled = true
	m.LastSession = session
}

func (m *MockEventHandler) OnTimerTick(session *Session, remainingTime time.Duration) {
	m.TimerTickCalled = true
	m.LastSession = session
	m.LastRemainingTime = remainingTime
}

func (m *MockEventHandler) OnWarning(idleDuration time.Duration) {
	m.WarningCalled = true
	m.LastIdleDuration = idleDuration
}

func TestPomodoroManager_NewPomodoroManager(t *testing.T) {
	config := NewDefaultConfig()
	handler := &MockEventHandler{}
	
	manager := NewPomodoroManager(config, handler)
	
	if manager.GetCurrentState() != SessionStateIdle {
		t.Errorf("Expected initial state to be SessionStateIdle, got %v", manager.GetCurrentState())
	}
	
	if manager.GetCurrentSession() != nil {
		t.Error("Expected initial session to be nil")
	}
}

func TestPomodoroManager_StartWorkSession(t *testing.T) {
	config := NewDefaultConfig()
	handler := &MockEventHandler{}
	manager := NewPomodoroManager(config, handler)
	
	err := manager.StartWorkSession()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	if manager.GetCurrentState() != SessionStateWork {
		t.Errorf("Expected state to be SessionStateWork, got %v", manager.GetCurrentState())
	}
	
	if !handler.SessionStartCalled {
		t.Error("Expected OnSessionStart to be called")
	}
	
	session := manager.GetCurrentSession()
	if session == nil {
		t.Error("Expected current session to not be nil")
	}
	
	if session.State != SessionStateWork {
		t.Errorf("Expected session state to be SessionStateWork, got %v", session.State)
	}
}

func TestPomodoroManager_StartBreakSession(t *testing.T) {
	config := NewDefaultConfig()
	handler := &MockEventHandler{}
	manager := NewPomodoroManager(config, handler)
	
	err := manager.StartBreakSession()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	if manager.GetCurrentState() != SessionStateBreak {
		t.Errorf("Expected state to be SessionStateBreak, got %v", manager.GetCurrentState())
	}
	
	if !handler.SessionStartCalled {
		t.Error("Expected OnSessionStart to be called")
	}
	
	session := manager.GetCurrentSession()
	if session == nil {
		t.Error("Expected current session to not be nil")
	}
	
	if session.State != SessionStateBreak {
		t.Errorf("Expected session state to be SessionStateBreak, got %v", session.State)
	}
}

func TestPomodoroManager_PauseResumeSession(t *testing.T) {
	config := NewDefaultConfig()
	handler := &MockEventHandler{}
	manager := NewPomodoroManager(config, handler)
	
	// Start a work session
	manager.StartWorkSession()
	
	// Pause the session
	err := manager.PauseSession()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	if !handler.SessionPauseCalled {
		t.Error("Expected OnSessionPause to be called")
	}
	
	session := manager.GetCurrentSession()
	if !session.IsPaused {
		t.Error("Expected session to be paused")
	}
	
	// Resume the session
	err = manager.ResumeSession()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	if !handler.SessionResumeCalled {
		t.Error("Expected OnSessionResume to be called")
	}
	
	if session.IsPaused {
		t.Error("Expected session to not be paused after resume")
	}
}

func TestPomodoroManager_GetRemainingTime(t *testing.T) {
	config := &SessionConfig{
		WorkDuration:  2 * time.Second,
		BreakDuration: 1 * time.Second,
	}
	handler := &MockEventHandler{}
	manager := NewPomodoroManager(config, handler)
	
	// No session initially
	if manager.GetRemainingTime() != 0 {
		t.Error("Expected remaining time to be 0 when no session")
	}
	
	// Start work session
	manager.StartWorkSession()
	
	remaining := manager.GetRemainingTime()
	if remaining <= 0 || remaining > config.WorkDuration {
		t.Errorf("Expected remaining time to be between 0 and %v, got %v", config.WorkDuration, remaining)
	}
}

func TestPomodoroManager_SessionExpiration(t *testing.T) {
	config := &SessionConfig{
		WorkDuration:  100 * time.Millisecond,
		BreakDuration: 50 * time.Millisecond,
	}
	handler := &MockEventHandler{}
	manager := NewPomodoroManager(config, handler)
	
	// Start work session
	manager.StartWorkSession()
	
	// Wait for session to expire
	time.Sleep(200 * time.Millisecond)
	
	// Session should be expired and state should be Idle
	if manager.GetCurrentState() != SessionStateIdle {
		t.Errorf("Expected state to be SessionStateIdle after session expiry, got %v", manager.GetCurrentState())
	}
	
	if !handler.SessionEndCalled {
		t.Error("Expected OnSessionEnd to be called")
	}
}

func TestPomodoroManager_GetIdleDuration(t *testing.T) {
	config := NewDefaultConfig()
	handler := &MockEventHandler{}
	manager := NewPomodoroManager(config, handler)
	
	// Should have some idle duration initially
	idleDuration := manager.GetIdleDuration()
	if idleDuration <= 0 {
		t.Error("Expected idle duration to be greater than 0")
	}
	
	// Start a session - idle duration should be 0
	manager.StartWorkSession()
	idleDuration = manager.GetIdleDuration()
	if idleDuration != 0 {
		t.Errorf("Expected idle duration to be 0 during session, got %v", idleDuration)
	}
}