package domain

import (
	"testing"
	"time"
)

func TestSessionConfig_NewDefaultConfig(t *testing.T) {
	cfg := NewDefaultConfig()
	
	if cfg.WorkDuration != 25*time.Minute {
		t.Errorf("Expected work duration to be 25 minutes, got %v", cfg.WorkDuration)
	}
	
	if cfg.BreakDuration != 5*time.Minute {
		t.Errorf("Expected break duration to be 5 minutes, got %v", cfg.BreakDuration)
	}
}

func TestSession_NewWorkSession(t *testing.T) {
	config := NewDefaultConfig()
	session := NewWorkSession(config)
	
	if session.State != SessionStateWork {
		t.Errorf("Expected state to be SessionStateWork, got %v", session.State)
	}
	
	if session.Duration != config.WorkDuration {
		t.Errorf("Expected duration to be %v, got %v", config.WorkDuration, session.Duration)
	}
	
	if session.IsPaused {
		t.Error("Expected session to not be paused initially")
	}
}

func TestSession_NewBreakSession(t *testing.T) {
	config := NewDefaultConfig()
	session := NewBreakSession(config)
	
	if session.State != SessionStateBreak {
		t.Errorf("Expected state to be SessionStateBreak, got %v", session.State)
	}
	
	if session.Duration != config.BreakDuration {
		t.Errorf("Expected duration to be %v, got %v", config.BreakDuration, session.Duration)
	}
}

func TestSession_GetRemainingTime(t *testing.T) {
	config := &SessionConfig{
		WorkDuration: 2 * time.Second,
		BreakDuration: 1 * time.Second,
	}
	
	session := NewWorkSession(config)
	
	// Should have close to full duration remaining
	remaining := session.GetRemainingTime()
	if remaining > config.WorkDuration || remaining < config.WorkDuration-100*time.Millisecond {
		t.Errorf("Expected remaining time to be close to %v, got %v", config.WorkDuration, remaining)
	}
	
	// Wait a bit and check again
	time.Sleep(100 * time.Millisecond)
	remaining = session.GetRemainingTime()
	if remaining > config.WorkDuration-50*time.Millisecond {
		t.Errorf("Expected remaining time to decrease, got %v", remaining)
	}
}

func TestSession_IsExpired(t *testing.T) {
	config := &SessionConfig{
		WorkDuration: 100 * time.Millisecond,
		BreakDuration: 50 * time.Millisecond,
	}
	
	session := NewWorkSession(config)
	
	if session.IsExpired() {
		t.Error("Expected session to not be expired initially")
	}
	
	// Wait for session to expire
	time.Sleep(150 * time.Millisecond)
	
	if !session.IsExpired() {
		t.Error("Expected session to be expired after duration")
	}
}

func TestSession_PauseResume(t *testing.T) {
	config := NewDefaultConfig()
	session := NewWorkSession(config)
	
	// Pause session
	session.Pause()
	if !session.IsPaused {
		t.Error("Expected session to be paused")
	}
	
	pausedTime := session.GetRemainingTime()
	time.Sleep(50 * time.Millisecond)
	
	// Time should not change while paused
	if session.GetRemainingTime() != pausedTime {
		t.Error("Expected remaining time to not change while paused")
	}
	
	// Resume session
	session.Resume()
	if session.IsPaused {
		t.Error("Expected session to not be paused after resume")
	}
}

func TestSessionState_String(t *testing.T) {
	tests := []struct {
		state    SessionState
		expected string
	}{
		{SessionStateIdle, "Idle"},
		{SessionStateWork, "WorkSession"},
		{SessionStateBreak, "BreakSession"},
	}
	
	for _, test := range tests {
		if test.state.String() != test.expected {
			t.Errorf("Expected %v.String() to be %s, got %s", test.state, test.expected, test.state.String())
		}
	}
}