package application

import (
	"karedoro/domain"
)

type SessionService struct {
	session *domain.Session
	eventCallbacks map[string][]func()
}

func NewSessionService() *SessionService {
	service := &SessionService{
		session: domain.NewSession(),
		eventCallbacks: make(map[string][]func()),
	}
	
	service.session.AddStateChangeCallback(service.onStateChange)
	
	return service
}

func (s *SessionService) StartWorkSession() error {
	err := s.session.StartWorkSession()
	if err != nil {
		return err
	}
	
	s.triggerEvent("work_session_start")
	return nil
}

func (s *SessionService) StartBreakSession() error {
	err := s.session.StartBreakSession()
	if err != nil {
		return err
	}
	
	s.triggerEvent("break_session_start")
	return nil
}

func (s *SessionService) PauseSession() error {
	err := s.session.PauseSession()
	if err != nil {
		return err
	}
	
	s.triggerEvent("session_pause")
	return nil
}

func (s *SessionService) ResumeSession() error {
	err := s.session.ResumeSession()
	if err != nil {
		return err
	}
	
	s.triggerEvent("session_resume")
	return nil
}

func (s *SessionService) Update() {
	shouldShowWarning := s.session.ShouldShowWarning()
	
	s.session.Update()
	
	// 待機状態で警告タイマーが終了したら、強制的に警告を発動
	if shouldShowWarning && s.session.GetState() == domain.Idle {
		s.triggerEvent("warning")
		s.session.ResetWarningTimer()
	}
}

func (s *SessionService) GetSession() *domain.Session {
	return s.session
}

func (s *SessionService) AddEventCallback(eventName string, callback func()) {
	if s.eventCallbacks[eventName] == nil {
		s.eventCallbacks[eventName] = make([]func(), 0)
	}
	s.eventCallbacks[eventName] = append(s.eventCallbacks[eventName], callback)
}

func (s *SessionService) triggerEvent(eventName string) {
	if callbacks, exists := s.eventCallbacks[eventName]; exists {
		for _, callback := range callbacks {
			callback()
		}
	}
}

func (s *SessionService) onStateChange(oldState, newState domain.SessionState) {
	switch newState {
	case domain.WorkSession:
		if oldState == domain.Idle {
			s.triggerEvent("work_session_start")
		}
	case domain.BreakSession:
		if oldState == domain.Idle {
			s.triggerEvent("break_session_start")
		}
	case domain.Idle:
		if oldState == domain.WorkSession {
			s.triggerEvent("work_session_end")
		} else if oldState == domain.BreakSession {
			s.triggerEvent("break_session_end")
		}
	}
}