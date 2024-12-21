package main

import (
	"auth/config"
	"auth/internal/api"
	"auth/internal/repository"
	"auth/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token"},
	}))

	defaultConfig := config.LoadDefaultConfig()

	db := ConnectToDB(defaultConfig)

	authRepo := repository.NewAuthRepository(db)
	sessionRepo := repository.NewSessionRepository(db)

	sessionService := service.NewSessionService(sessionRepo)
	tokenService := service.NewTokenService()
	authService := service.NewAuthService(authRepo, sessionService, tokenService)

	authAPI := api.NewAuthAPI(authService, sessionService, tokenService)

	authGroup := r.Group("/")
	authAPI.RegisterRoutes(authGroup)

	err := r.Run(":8080")

	if err != nil {
		return
	}
}

func ConnectToDB(config *config.Config) *gorm.DB {
	dsn := "host=" + config.Database.Host + " user=" + config.Database.User +
		" password=" + config.Database.Password +
		" dbname=" + config.Database.Name +
		" port=" + strconv.Itoa(config.Database.Port) +
		" sslmode=disable TimeZone=Asia/Aqtobe"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
