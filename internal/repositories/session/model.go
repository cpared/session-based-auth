package session

import "time"

type Session struct {
	ID             string
	CreationDate   *time.Time
	ExpirationDate *time.Time
	TTL            int
	Auth           Auth
}

type Auth struct {
	Role string
}
