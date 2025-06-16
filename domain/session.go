package domain

import (
	"time"
)

type Session struct {
	state            SessionState
	currentTimer     *Timer
	warningTimer     *Timer
	sessionType      SessionType
	lastWarningTime  time.Time
	stateChangeCallbacks []func(SessionState, SessionState)
}

func NewSession() *Session {
	return &Session{
		state:                Idle,
		currentTimer:         NewTimer(0),
		warningTimer:         NewTimer(WarningInterval),
		sessionType:          Work,
		stateChangeCallbacks: make([]func(SessionState, SessionState), 0),
	}
}

func (s *Session) AddStateChangeCallback(callback func(SessionState, SessionState)) {
	s.stateChangeCallbacks = append(s.stateChangeCallbacks, callback)
}

func (s *Session) setState(newState SessionState) {
	oldState := s.state
	s.state = newState
	
	for _, callback := range s.stateChangeCallbacks {
		callback(oldState, newState)
	}
}

func (s *Session) StartWorkSession() error {
	if s.state != Idle {
		return nil
	}
	
	s.sessionType = Work
	s.currentTimer.Reset(WorkSessionDuration)
	s.currentTimer.Start()
	s.warningTimer.Stop()
	s.setState(WorkSession)
	
	return nil
}

func (s *Session) StartBreakSession() error {
	if s.state != Idle {
		return nil
	}
	
	s.sessionType = Break
	s.currentTimer.Reset(BreakSessionDuration)
	s.currentTimer.Start()
	s.warningTimer.Stop()
	s.setState(BreakSession)
	
	return nil
}

func (s *Session) PauseSession() error {
	if s.state != WorkSession && s.state != BreakSession {
		return nil
	}
	
	s.currentTimer.Pause()
	return nil
}

func (s *Session) ResumeSession() error {
	if s.state != WorkSession && s.state != BreakSession {
		return nil
	}
	
	s.currentTimer.Resume()
	return nil
}

func (s *Session) Update() {
	s.currentTimer.Update()
	
	if s.currentTimer.IsFinished() {
		switch s.state {
		case WorkSession:
			s.setState(Idle)
			s.warningTimer.Reset(WarningInterval)
			s.warningTimer.Start()
		case BreakSession:
			s.setState(Idle)
			s.warningTimer.Reset(WarningInterval)
			s.warningTimer.Start()
		}
	}
	
	if s.state == Idle {
		s.warningTimer.Update()
	}
}

func (s *Session) ShouldShowWarning() bool {
	if s.state != Idle {
		return false
	}
	
	return s.warningTimer.IsFinished()
}

func (s *Session) ResetWarningTimer() {
	s.warningTimer.Reset(WarningInterval)
	s.warningTimer.Start()
}

func (s *Session) GetState() SessionState {
	return s.state
}

func (s *Session) GetSessionType() SessionType {
	return s.sessionType
}

func (s *Session) GetCurrentTimer() *Timer {
	return s.currentTimer
}

func (s *Session) GetWarningTimer() *Timer {
	return s.warningTimer
}

func (s *Session) IsSessionActive() bool {
	return s.state == WorkSession || s.state == BreakSession
}

func (s *Session) IsSessionPaused() bool {
	return s.currentTimer.IsPaused()
}

func (s *Session) GetTimeRemaining() time.Duration {
	return s.currentTimer.Remaining()
}

func (s *Session) GetProgress() float64 {
	return s.currentTimer.Progress()
}