package service

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID uuid.UUID
	CreationDate *time.Time
	ExpirationDate *time.Time
}

type Service struct {
	sessions map[string]*Session
}

func NewService() *Service {
	return &Service{
		sessions: map[string]*Session{},
	}
}

func (s *Service) Get(userID string) uuid.UUID {
	val, found := s.sessions[userID]
	if !found {
		return uuid.Nil
	}

	if val.ExpirationDate.Before(time.Now()) {
		delete(s.sessions, userID)
		return uuid.Nil
	}
	
	return val.ID
}

func (s *Service) Save(userID string, sess uuid.UUID) {
	t := time.Now()
	exp := t.Add(30 * time.Minute)
	s.sessions[userID] = &Session{
		ID: sess,
		CreationDate: &t,
		ExpirationDate: &exp,
	}
}