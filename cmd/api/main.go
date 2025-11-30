package main

import (
	pokehdl "session-based-auth/internal/handlers/pokemon"
	sesshdl "session-based-auth/internal/handlers/session_based"
	pokerepo "session-based-auth/internal/repositories/pokemon"
	sesssvc "session-based-auth/internal/services"
	pokesvc "session-based-auth/internal/services/pokemon"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Init Repositories
	pokeRepo := pokerepo.New()

	// Init Services
	pokeSvc := pokesvc.New(pokeRepo)
	sessSvc := sesssvc.New()

	// Init Handlers
	sess := sesshdl.New(sessSvc)
	pokeHdl := pokehdl.New(pokeSvc)

	// Routes
	r.POST("/login", sess.Validate())
	r.GET("/types/pokemons/:id", pokeHdl.GetPokemonData())

	r.Run()
}
