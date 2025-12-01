package session

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const (
	Admin = "ADMIN"
	TTL   = 30 //mins
)

type Repository struct {
	validUsers map[string]string
	sessions   map[string]*Session
}

func New() *Repository {
	return &Repository{
		validUsers: map[string]string{
			"cpared": "12345",
		},
		sessions: map[string]*Session{},
	}
}

func (r *Repository) Create(ctx context.Context, user, password string) *Session {
	pass, found := r.validUsers[user]
	if !found || pass != password {
		return &Session{}
	}

	now := time.Now()
	exp := now.Add(TTL * time.Minute)
	sessID := uuid.NewString()

	sess := &Session{
		ID:             sessID,
		CreationDate:   &now,
		ExpirationDate: &exp,
		TTL:            TTL,
		Auth: Auth{
			Role: Admin,
		},
	}

	r.sessions[sessID] = sess
	return sess
}

func (r *Repository) Get(ctx context.Context, sessionID string) *Session {
	if sess, found := r.sessions[sessionID]; found {
		return sess
	}
	return &Session{}
}

func (r *Repository) Delete(ctx context.Context, sessionID string) *Session {
	if sess, found := r.sessions[sessionID]; found {
		delete(r.sessions, sessionID)
		return sess
	}
	return &Session{}
}

func (r *Repository) Refresh(ctx context.Context, sessionID string) *Session {
	if sess, found := r.sessions[sessionID]; found {
		exp := time.Now().Add(TTL * time.Minute)
		sess.ExpirationDate = &exp
		return sess
	}
	return &Session{}
}
