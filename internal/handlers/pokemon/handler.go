package handlers

import (
	"net/http"
	repositories "session-based-auth/internal/repositories/pokemon"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetPokemonDataByID(id string) *repositories.Type
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
		// 1. Validate auth
		id := c.Param("id")

		// 2. Call service
		data := h.service.GetPokemonDataByID(id)

		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
