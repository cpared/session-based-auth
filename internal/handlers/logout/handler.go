package logut

import (
	"context"
	"net/http"
	"session-based-auth/internal/repositories/session"

	"github.com/gin-gonic/gin"
)

const (
	SessionCookieName = "sessionID"
)

type Service interface {
	Delete(ctx context.Context, sessID string) string
}

type Handler struct{
	service Service
}

func New(serv Service) *Handler {
	return &Handler{
		service: serv,
	}
}

func (h *Handler) Logout() gin.HandlerFunc{
	return func(c *gin.Context) {
		sessValue, found := c.Get("session")
		if !found {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid session"})
			return
		}

		sess, ok := sessValue.(*session.Session)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot cast session"})
			return
		}

		h.service.Delete(c.Request.Context(), sess.ID)

		c.JSON(http.StatusOK, gin.H{"message": "logout succesfull!"})
	}
}