package sessionbased

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var validUsers = map[string]string{
	"cpared": "12345",
}

type Service interface {
	Get(userID string) string
	Save(userID string, sess string)
}

type Body struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type Handler struct {
	Service Service
}

func New(serv Service) *Handler {
	return &Handler{
		Service: serv,
	}
}

func (h *Handler) Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body Body
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 1. Authentication
		pass, found := validUsers[body.User]
		if !found || pass != body.Password {
			fmt.Printf("Invalid credentials")
			c.JSON(http.StatusUnauthorized, gin.H{"err": "invalid credentials"})
			return
		}

		if sessionID := h.Service.Get(body.User); sessionID != "" {
			fmt.Println("session token already exist!")
			c.JSON(http.StatusOK, gin.H{"message": "login OK"})
			return
		}

		// 2. Create UUID for session
		sessionID := uuid.NewString()
		c.SetCookie(
			"sessionID",
			sessionID,
			3600, // Duration
			"/",
			"localhost",
			false, // Secure: only HTTPS
			true,  // HttpOnly: no accesible from JS
		)

		h.Service.Save(body.User, sessionID)
		c.JSON(http.StatusOK, gin.H{"message": "login OK"})
	}
}
