package handlers

import (
	"context"
	"net/http"
	repositories "session-based-auth/internal/repositories/pokemon"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetPokemonDataByID(ctx context.Context,id string) *repositories.Type
}

type Handler struct {
	service Service
}

func New(serv Service) *Handler {
	return &Handler{
		service: serv,
	}
}

func (h *Handler) GetPokemonData() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid pokemon id or name"})
			return
		}

		data := h.service.GetPokemonDataByID(c.Request.Context(), id)

		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
