package main

import (
	pokehdl "session-based-auth/internal/handlers/pokemon"
	loginhdl "session-based-auth/internal/handlers/login"
	logouthdl "session-based-auth/internal/handlers/logout"
	middleware "session-based-auth/internal/middleware"
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
	loginHdl := loginhdl.New(sessSvc)
	logoutHdl := logouthdl.New(sessSvc)
	pokeHdl := pokehdl.New(pokeSvc)

	// Routes
	r.POST("/login", loginHdl.Login())
	r.POST("/logout", middleware.Auth(sessSvc), logoutHdl.Logout())
	r.GET("/types/pokemons/:id", middleware.Auth(sessSvc), pokeHdl.GetPokemonData())

	r.Run()
}
