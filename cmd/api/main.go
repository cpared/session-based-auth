package main

import (
	sessionbased "session-based-auth/internal/handlers/session_based"
	serv "session-based-auth/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	service := serv.NewService()
	sess := sessionbased.NewHandler(service)

	r.POST("/login", sess.Validate())

  	r.Run()
}