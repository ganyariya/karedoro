package domain

import (
	"testing"
	"time"
)

func TestSession_InitialState(t *testing.T) {
	session := NewSession()
	
	if session.GetState() != Idle {
		t.Errorf("Expected initial state to be Idle, got %v", session.GetState())
	}
	
	if session.IsSessionActive() {
		t.Error("Session should not be active initially")
	}
	
	if session.IsSessionPaused() {
		t.Error("Session should not be paused initially")
	}
}

func TestSession_StartWorkSession(t *testing.T) {
	session := NewSession()
	
	err := session.StartWorkSession()
	if err != nil {
		t.Errorf("StartWorkSession should not return error, got %v", err)
	}
	
	if session.GetState() != WorkSession {
		t.Errorf("Expected state to be WorkSession, got %v", session.GetState())
	}
	
	if session.GetSessionType() != Work {
		t.Errorf("Expected session type to be Work, got %v", session.GetSessionType())
	}
	
	if !session.IsSessionActive() {
		t.Error("Session should be active after starting work session")
	}
	
	// Check timer is properly set up
	remaining := session.GetTimeRemaining()
	if remaining > WorkSessionDuration || remaining < WorkSessionDuration-time.Second {
		t.Errorf("Expected remaining time around %v, got %v", WorkSessionDuration, remaining)
	}
}

func TestSession_StartBreakSession(t *testing.T) {
	session := NewSession()
	
	err := session.StartBreakSession()
	if err != nil {
		t.Errorf("StartBreakSession should not return error, got %v", err)
	}
	
	if session.GetState() != BreakSession {
		t.Errorf("Expected state to be BreakSession, got %v", session.GetState())
	}
	
	if session.GetSessionType() != Break {
		t.Errorf("Expected session type to be Break, got %v", session.GetSessionType())
	}
	
	if !session.IsSessionActive() {
		t.Error("Session should be active after starting break session")
	}
	
	// Check timer is properly set up
	remaining := session.GetTimeRemaining()
	if remaining > BreakSessionDuration || remaining < BreakSessionDuration-time.Second {
		t.Errorf("Expected remaining time around %v, got %v", BreakSessionDuration, remaining)
	}
}

func TestSession_PauseAndResume(t *testing.T) {
	session := NewSession()
	
	// Start work session
	session.StartWorkSession()
	
	// Wait a bit
	time.Sleep(100 * time.Millisecond)
	
	// Pause
	err := session.PauseSession()
	if err != nil {
		t.Errorf("PauseSession should not return error, got %v", err)
	}
	
	if !session.IsSessionPaused() {
		t.Error("Session should be paused after PauseSession")
	}
	
	if session.GetState() != WorkSession {
		t.Errorf("State should remain WorkSession when paused, got %v", session.GetState())
	}
	
	// Resume
	err = session.ResumeSession()
	if err != nil {
		t.Errorf("ResumeSession should not return error, got %v", err)
	}
	
	if session.IsSessionPaused() {
		t.Error("Session should not be paused after ResumeSession")
	}
	
	if !session.IsSessionActive() {
		t.Error("Session should be active after ResumeSession")
	}
}

func TestSession_CannotStartSessionWhenActive(t *testing.T) {
	session := NewSession()
	
	// Start work session
	session.StartWorkSession()
	
	// Try to start another work session (should be ignored)
	session.StartWorkSession()
	if session.GetState() != WorkSession {
		t.Error("State should remain WorkSession")
	}
	
	// Try to start break session (should be ignored)
	session.StartBreakSession()
	if session.GetState() != WorkSession {
		t.Error("State should remain WorkSession")
	}
}

func TestSession_StateTransitionOnCompletion(t *testing.T) {
	session := NewSession()
	
	// Use a very short duration for testing
	session.currentTimer.Reset(50 * time.Millisecond)
	session.currentTimer.Start()
	session.state = WorkSession
	
	// Wait for timer to complete
	time.Sleep(100 * time.Millisecond)
	
	// Update should trigger state change
	session.Update()
	
	if session.GetState() != Idle {
		t.Errorf("Expected state to be Idle after timer completion, got %v", session.GetState())
	}
	
	if session.IsSessionActive() {
		t.Error("Session should not be active after completion")
	}
}

func TestSession_WarningSystem(t *testing.T) {
	session := NewSession()
	
	// Manually set to idle state with short warning interval
	session.setState(Idle)
	session.warningTimer.Reset(50 * time.Millisecond)
	session.warningTimer.Start()
	
	// Initially should not show warning
	if session.ShouldShowWarning() {
		t.Error("Should not show warning initially")
	}
	
	// Wait for warning timer to complete
	time.Sleep(100 * time.Millisecond)
	session.warningTimer.Update()
	
	if !session.ShouldShowWarning() {
		t.Error("Should show warning after timer expires")
	}
	
	// Reset warning timer
	session.ResetWarningTimer()
	
	if session.ShouldShowWarning() {
		t.Error("Should not show warning after reset")
	}
}

func TestSession_StateChangeCallback(t *testing.T) {
	session := NewSession()
	
	var callbackCalled bool
	var oldState, newState SessionState
	
	// Add callback
	session.AddStateChangeCallback(func(old, new SessionState) {
		callbackCalled = true
		oldState = old
		newState = new
	})
	
	// Start work session (should trigger callback)
	session.StartWorkSession()
	
	if !callbackCalled {
		t.Error("State change callback should have been called")
	}
	
	if oldState != Idle {
		t.Errorf("Expected old state to be Idle, got %v", oldState)
	}
	
	if newState != WorkSession {
		t.Errorf("Expected new state to be WorkSession, got %v", newState)
	}
}

func TestSession_MultipleCallbacks(t *testing.T) {
	session := NewSession()
	
	var callback1Called, callback2Called bool
	
	// Add multiple callbacks
	session.AddStateChangeCallback(func(old, new SessionState) {
		callback1Called = true
	})
	
	session.AddStateChangeCallback(func(old, new SessionState) {
		callback2Called = true
	})
	
	// Trigger state change
	session.StartWorkSession()
	
	if !callback1Called {
		t.Error("First callback should have been called")
	}
	
	if !callback2Called {
		t.Error("Second callback should have been called")
	}
}

func TestSession_CannotPauseIdleSession(t *testing.T) {
	session := NewSession()
	
	// Try to pause idle session
	err := session.PauseSession()
	if err != nil {
		t.Errorf("PauseSession should not return error even for idle session, got %v", err)
	}
	
	// State should remain idle
	if session.GetState() != Idle {
		t.Errorf("Expected state to remain Idle, got %v", session.GetState())
	}
}

func TestSession_CannotResumeNonPausedSession(t *testing.T) {
	session := NewSession()
	
	// Start work session
	session.StartWorkSession()
	
	// Try to resume already running session
	err := session.ResumeSession()
	if err != nil {
		t.Errorf("ResumeSession should not return error, got %v", err)
	}
	
	// Should still be running
	if !session.IsSessionActive() {
		t.Error("Session should still be active")
	}
}

func TestSession_GetProgress(t *testing.T) {
	session := NewSession()
	
	// Start work session
	session.StartWorkSession()
	
	// Progress should be very small initially
	progress := session.GetProgress()
	if progress < 0 || progress > 0.1 {
		t.Errorf("Expected progress between 0 and 0.1, got %v", progress)
	}
	
	// Wait a bit
	time.Sleep(100 * time.Millisecond)
	
	// Progress should have increased
	newProgress := session.GetProgress()
	if newProgress <= progress {
		t.Errorf("Progress should have increased from %v to %v", progress, newProgress)
	}
}

func TestSessionType_Duration(t *testing.T) {
	workDuration := Work.Duration()
	if workDuration != WorkSessionDuration {
		t.Errorf("Expected work duration %v, got %v", WorkSessionDuration, workDuration)
	}
	
	breakDuration := Break.Duration()
	if breakDuration != BreakSessionDuration {
		t.Errorf("Expected break duration %v, got %v", BreakSessionDuration, breakDuration)
	}
}

func TestSessionType_String(t *testing.T) {
	if Work.String() != "Work" {
		t.Errorf("Expected Work.String() to be 'Work', got %v", Work.String())
	}
	
	if Break.String() != "Break" {
		t.Errorf("Expected Break.String() to be 'Break', got %v", Break.String())
	}
}

func TestSessionState_String(t *testing.T) {
	if WorkSession.String() != "WorkSession" {
		t.Errorf("Expected WorkSession.String() to be 'WorkSession', got %v", WorkSession.String())
	}
	
	if BreakSession.String() != "BreakSession" {
		t.Errorf("Expected BreakSession.String() to be 'BreakSession', got %v", BreakSession.String())
	}
	
	if Idle.String() != "Idle" {
		t.Errorf("Expected Idle.String() to be 'Idle', got %v", Idle.String())
	}
}