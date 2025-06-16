package domain

import (
	"testing"
	"time"
)

func TestTimer_PauseAndResume(t *testing.T) {
	timer := NewTimer(5 * time.Second)
	
	// Start timer
	timer.Start()
	
	// Wait 1 second
	time.Sleep(1 * time.Second)
	
	// Pause timer
	timer.Pause()
	
	// Check remaining time (should be around 4 seconds)
	remaining := timer.Remaining()
	if remaining > 4*time.Second || remaining < 3800*time.Millisecond {
		t.Errorf("Expected remaining time around 4s, got %v", remaining)
	}
	
	// Wait another second while paused
	time.Sleep(1 * time.Second)
	
	// Remaining time should not have changed
	stillRemaining := timer.Remaining()
	if stillRemaining != remaining {
		t.Errorf("Timer continued while paused. Expected %v, got %v", remaining, stillRemaining)
	}
	
	// Resume timer
	timer.Resume()
	
	// Wait another second
	time.Sleep(1 * time.Second)
	
	// Check remaining time (should be around 3 seconds)
	finalRemaining := timer.Remaining()
	if finalRemaining > 3*time.Second || finalRemaining < 2800*time.Millisecond {
		t.Errorf("Expected remaining time around 3s after resume, got %v", finalRemaining)
	}
}

func TestTimer_Reset(t *testing.T) {
	timer := NewTimer(5 * time.Second)
	
	// Start and wait
	timer.Start()
	time.Sleep(1 * time.Second)
	
	// Reset with new duration
	timer.Reset(10 * time.Second)
	
	// Check that timer is reset
	if timer.IsRunning() {
		t.Error("Timer should not be running after reset")
	}
	
	if timer.Remaining() != 10*time.Second {
		t.Errorf("Expected 10s after reset, got %v", timer.Remaining())
	}
}

func TestTimer_Finish(t *testing.T) {
	timer := NewTimer(100 * time.Millisecond)
	
	timer.Start()
	
	// Wait for timer to finish
	time.Sleep(200 * time.Millisecond)
	timer.Update()
	
	if !timer.IsFinished() {
		t.Error("Timer should be finished")
	}
	
	if timer.Remaining() != 0 {
		t.Errorf("Expected 0 remaining time, got %v", timer.Remaining())
	}
}