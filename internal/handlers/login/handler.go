package session

import (
	"context"
	"fmt"
	"net/http"
	"session-based-auth/internal/repositories/session"

	"github.com/gin-gonic/gin"
)

type Service interface {
	Get(ctx context.Context, userID string) *session.Session
	Create(ctx context.Context, user, password string) *session.Session
}

type Handler struct {
	service Service
}

func New(serv Service) *Handler {
	return &Handler{
		service: serv,
	}
}

func (h *Handler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body BodyRequest
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 1. Authentication
		sess := h.service.Create(c.Request.Context(), body.User, body.Password)
		if sess.ID == "" {
			fmt.Printf("Invalid credentials")
			c.JSON(http.StatusUnauthorized, gin.H{"err": "invalid credentials"})
			return
		}

		// 2. Set UUID cookie
		c.SetCookie(
			"sessionID",
			sess.ID,
			sess.TTL, // Duration
			"/",
			"localhost",
			false, // Secure: only HTTPS
			true,  // HttpOnly: no accesible from JS
		)

		c.JSON(http.StatusOK, gin.H{"message": "login OK"})
	}
}
