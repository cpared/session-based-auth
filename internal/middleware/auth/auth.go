package middleware

import (
	"context"
	"net/http"
	"session-based-auth/internal/repositories/session"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	SessionCookieName = "sessionID"
)

type Service interface {
	Get(ctx context.Context, userID string) *session.Session
	Create(ctx context.Context, user, password string) *session.Session
}

func Validate(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie(SessionCookieName)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"err": "Cookie is not found"})
			c.Abort()
			return 
		}

		sess := s.Get(c.Request.Context() ,cookie)
		if sess.ID == "" || sess.ExpirationDate.Before(time.Now()) {
			c.JSON(http.StatusUnauthorized, gin.H{"err": "invalid credentials"})
			c.Abort()
			return
		}

		c.Next()
	}
}