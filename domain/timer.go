package domain

import (
	"time"
)

type Timer struct {
	duration    time.Duration
	remaining   time.Duration
	isRunning   bool
	isPaused    bool
	startTime   time.Time
	pausedTime  time.Time
}

func NewTimer(duration time.Duration) *Timer {
	return &Timer{
		duration:  duration,
		remaining: duration,
		isRunning: false,
		isPaused:  false,
	}
}

func (t *Timer) Start() {
	if t.isPaused {
		t.Resume()
		return
	}
	
	t.isRunning = true
	t.isPaused = false
	t.startTime = time.Now()
}

func (t *Timer) Pause() {
	if !t.isRunning || t.isPaused {
		return
	}
	
	t.isPaused = true
	t.pausedTime = time.Now()
	t.remaining = t.remaining - time.Since(t.startTime)
}

func (t *Timer) Resume() {
	if !t.isPaused {
		return
	}
	
	t.isPaused = false
	t.startTime = time.Now()
}

func (t *Timer) Stop() {
	t.isRunning = false
	t.isPaused = false
	t.remaining = t.duration
}

func (t *Timer) Reset(duration time.Duration) {
	t.duration = duration
	t.remaining = duration
	t.isRunning = false
	t.isPaused = false
}

func (t *Timer) Update() {
	if !t.isRunning || t.isPaused {
		return
	}
	
	elapsed := time.Since(t.startTime)
	t.remaining = t.duration - elapsed
	
	if t.remaining <= 0 {
		t.remaining = 0
		t.isRunning = false
	}
}

func (t *Timer) IsFinished() bool {
	return t.remaining <= 0 && !t.isPaused
}

func (t *Timer) IsRunning() bool {
	return t.isRunning && !t.isPaused
}

func (t *Timer) IsPaused() bool {
	return t.isPaused
}

func (t *Timer) Remaining() time.Duration {
	if t.isPaused {
		return t.remaining
	}
	
	if !t.isRunning {
		return t.duration
	}
	
	elapsed := time.Since(t.startTime)
	remaining := t.duration - elapsed
	
	if remaining < 0 {
		return 0
	}
	
	return remaining
}

func (t *Timer) Progress() float64 {
	if t.duration == 0 {
		return 1.0
	}
	
	remaining := t.Remaining()
	return 1.0 - (float64(remaining) / float64(t.duration))
}