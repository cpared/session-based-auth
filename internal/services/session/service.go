package service

import (
	"context"
	"session-based-auth/internal/repositories/session"
	"time"
)

type SessionRepository interface{
	Create(ctx context.Context, user, password string) *session.Session
	Get(ctx context.Context, sessionID string) *session.Session
	Delete(ctx context.Context, sessionID string) *session.Session
	Refresh(ctx context.Context, sessionID string) *session.Session
}

type Service struct {
	repository SessionRepository
}

func New(r SessionRepository) *Service {
	return &Service{
		repository: r,
	}
}

func (s *Service) Get(ctx context.Context, userID string) *session.Session {
	sess := s.repository.Get(ctx, userID)
	if sess.ExpirationDate.Before(time.Now()) {
		return s.repository.Delete(ctx, sess.ID)
	}

	return sess
}

func (s *Service) Create(ctx context.Context, user, password string) *session.Session {
	return s.repository.Create(ctx, user, password)
}

func (s *Service) Delete(ctx context.Context, sessID string) string {
	return s.repository.Delete(ctx, sessID).ID
}