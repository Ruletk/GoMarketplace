package main

import (
	"auth/internal/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token"},
	}))

	authAPI := api.NewAuthAPI()

	authGroup := r.Group("/")
	authAPI.RegisterRoutes(authGroup)

	err := r.Run(":8080")

	if err != nil {
		return
	}
}
