package service

import (
	"time"
)

type Session struct {
	ID             string
	CreationDate   *time.Time
	ExpirationDate *time.Time
}

type Service struct {
	sessions map[string]*Session
}

func New() *Service {
	return &Service{
		sessions: map[string]*Session{},
	}
}

func (s *Service) Get(userID string) string {
	val, found := s.sessions[userID]
	if !found {
		return ""
	}

	if val.ExpirationDate.Before(time.Now()) {
		delete(s.sessions, userID)
		return ""
	}

	return val.ID
}

func (s *Service) Save(userID string, sess string) {
	t := time.Now()
	exp := t.Add(30 * time.Minute)
	s.sessions[userID] = &Session{
		ID:             sess,
		CreationDate:   &t,
		ExpirationDate: &exp,
	}
}
