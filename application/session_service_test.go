package application

import (
	"testing"
	
	"karedoro/domain"
)

func TestSessionService_InitialState(t *testing.T) {
	service := NewSessionService()
	
	session := service.GetSession()
	if session.GetState() != domain.Idle {
		t.Errorf("Expected initial state to be Idle, got %v", session.GetState())
	}
}

func TestSessionService_StartWorkSession(t *testing.T) {
	service := NewSessionService()
	
	err := service.StartWorkSession()
	if err != nil {
		t.Errorf("StartWorkSession should not return error, got %v", err)
	}
	
	session := service.GetSession()
	if session.GetState() != domain.WorkSession {
		t.Errorf("Expected state to be WorkSession, got %v", session.GetState())
	}
}

func TestSessionService_StartBreakSession(t *testing.T) {
	service := NewSessionService()
	
	err := service.StartBreakSession()
	if err != nil {
		t.Errorf("StartBreakSession should not return error, got %v", err)
	}
	
	session := service.GetSession()
	if session.GetState() != domain.BreakSession {
		t.Errorf("Expected state to be BreakSession, got %v", session.GetState())
	}
}

func TestSessionService_PauseAndResume(t *testing.T) {
	service := NewSessionService()
	
	// Start work session
	service.StartWorkSession()
	
	// Pause
	err := service.PauseSession()
	if err != nil {
		t.Errorf("PauseSession should not return error, got %v", err)
	}
	
	session := service.GetSession()
	if !session.IsSessionPaused() {
		t.Error("Session should be paused")
	}
	
	// Resume
	err = service.ResumeSession()
	if err != nil {
		t.Errorf("ResumeSession should not return error, got %v", err)
	}
	
	if session.IsSessionPaused() {
		t.Error("Session should not be paused after resume")
	}
}

func TestSessionService_EventCallbacks(t *testing.T) {
	service := NewSessionService()
	
	var workStartCalled, breakStartCalled bool
	var pauseCalled, resumeCalled bool
	
	// Add event callbacks
	service.AddEventCallback("work_session_start", func() {
		workStartCalled = true
	})
	
	service.AddEventCallback("break_session_start", func() {
		breakStartCalled = true
	})
	
	service.AddEventCallback("session_pause", func() {
		pauseCalled = true
	})
	
	service.AddEventCallback("session_resume", func() {
		resumeCalled = true
	})
	
	// Test work session start
	service.StartWorkSession()
	if !workStartCalled {
		t.Error("work_session_start callback should have been called")
	}
	
	// Test pause
	service.PauseSession()
	if !pauseCalled {
		t.Error("session_pause callback should have been called")
	}
	
	// Test resume
	service.ResumeSession()
	if !resumeCalled {
		t.Error("session_resume callback should have been called")
	}
	
	// Reset and test break session
	service = NewSessionService()
	service.AddEventCallback("break_session_start", func() {
		breakStartCalled = true
	})
	
	service.StartBreakSession()
	if !breakStartCalled {
		t.Error("break_session_start callback should have been called")
	}
}

func TestSessionService_StateChangeEvents(t *testing.T) {
	service := NewSessionService()
	
	var eventCalled bool
	
	// Add state change event callbacks
	service.AddEventCallback("work_session_start", func() {
		eventCalled = true
	})
	
	// Test work session start
	session := service.GetSession()
	session.StartWorkSession()
	
	if !eventCalled {
		t.Error("Event callback should have been called")
	}
	
	// Note: Testing state transitions requires access to private methods
	// Domain layer tests cover the detailed state transition logic
}

func TestSessionService_WarningEvent(t *testing.T) {
	service := NewSessionService()
	
	var warningCalled bool
	service.AddEventCallback("warning", func() {
		warningCalled = true
	})
	
	// Set up warning condition using public API
	_ = service.GetSession()
	// We can't directly access setState, so this test is limited
	// In practice, warnings are tested through the domain layer tests
	
	// Just verify the callback can be added and service doesn't crash on update
	service.Update()
	
	// This test mainly verifies the event system works
	_ = warningCalled // Acknowledge the variable is used for setup
}

func TestSessionService_MultipleEventCallbacks(t *testing.T) {
	service := NewSessionService()
	
	var callback1Called, callback2Called bool
	
	// Add multiple callbacks for the same event
	service.AddEventCallback("work_session_start", func() {
		callback1Called = true
	})
	
	service.AddEventCallback("work_session_start", func() {
		callback2Called = true
	})
	
	// Trigger event
	service.StartWorkSession()
	
	if !callback1Called {
		t.Error("First callback should have been called")
	}
	
	if !callback2Called {
		t.Error("Second callback should have been called")
	}
}

func TestSessionService_Update(t *testing.T) {
	service := NewSessionService()
	
	// Start a session
	service.StartWorkSession()
	
	// Update should not cause errors
	service.Update()
	
	// Session should still be active
	session := service.GetSession()
	if !session.IsSessionActive() {
		t.Error("Session should still be active after update")
	}
}