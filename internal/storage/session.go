package storage

import (
	"fmt"

	"github.com/reyesossorio/f1-terminal/internal/domain"
)

// key->sessionkey
type SessionStorage struct {
	sessions map[int]*domain.SessionResult
	curSession int 
}

func NewSessionStorage() *SessionStorage {
	return &SessionStorage{
		sessions: make(map[int]*domain.SessionResult),
		curSession: -1,
	}
}

func (s *SessionStorage) SaveSession(session domain.Session) error {
	_, ok := s.sessions[session.SessionKey]

	if ok {
		return fmt.Errorf("session already stored")
	}

	s.sessions[session.SessionKey] = &domain.SessionResult{
		SessionName: session.SessionName,
		SessionType: session.SessionType,
	}
	return nil
}

func (s *SessionStorage) GetSessionInfo(sessionKey int) (*domain.SessionResult, error){
	val, ok := s.sessions[sessionKey]

	if !ok{
		return &domain.SessionResult{}, fmt.Errorf("session with key %d not found", sessionKey)
	}

	return val, nil
}

func (s *SessionStorage) SetCurSession (sessionKey int) {
	s.curSession = sessionKey
}

func (s *SessionStorage) GetCurSession () int {
	return s.curSession
}
