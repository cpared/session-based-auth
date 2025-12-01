package main

import (
	pokehdl "session-based-auth/internal/handlers/pokemon"
	sesshdl "session-based-auth/internal/handlers/session"
	auth "session-based-auth/internal/middleware/auth"
	pokerepo "session-based-auth/internal/repositories/pokemon"
	sessrepo "session-based-auth/internal/repositories/session"
	pokesvc "session-based-auth/internal/services/pokemon"
	sesssvc "session-based-auth/internal/services/session"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Global auth
	// r.Use(AuthMiddleware(repo))

	// Init Repositories
	pokeRepo := pokerepo.New()
	sessRepo := sessrepo.New()

	// Init Services
	pokeSvc := pokesvc.New(pokeRepo)
	sessSvc := sesssvc.New(sessRepo)

	// Init Handlers
	sess := sesshdl.New(sessSvc)
	pokeHdl := pokehdl.New(pokeSvc)

	// Routes
	r.POST("/login", sess.Login())
	// r.POST("/logut", gin.H{"message": "logued out!"})
	r.GET("/types/pokemons/:id", auth.Validate(sessSvc) ,pokeHdl.GetPokemonData())

	r.Run()
}
