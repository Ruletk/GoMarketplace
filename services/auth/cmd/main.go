package main

import (
	"auth/config"
	"auth/internal/api"
	"auth/internal/repository"
	"auth/internal/service"
	"auth/pkg/auth"
	"github.com/Ruletk/GoMarketplace/pkg/logging"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

func main() {
	logging.InitLogger(logging.LogConfig{
		Format: "json",
		Level:  "debug",
	})

	logging.Logger.Info("Starting the server")

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token", "Cookie"},
		AllowCredentials: true,
	}))

	defaultConfig := config.LoadDefaultConfig()

	db := ConnectToDB(defaultConfig)

	authRepo := repository.NewAuthRepository(db)
	sessionRepo := repository.NewSessionRepository(db)

	sessionService := service.NewSessionService(sessionRepo)
	tokenService := service.NewTokenService()
	authService := service.NewAuthService(authRepo, sessionService, tokenService)

	authAPI := api.NewAuthAPI(authService, sessionService, tokenService)

	public := r.Group("/")
	authAPI.RegisterPublicRoutes(public)

	unAuth := r.Group("/")
	unAuth.Use(auth.NoAuthMiddleware())
	authAPI.RegisterPublicOnlyRoutes(unAuth)

	private := r.Group("/")
	private.Use(auth.CookieTokenMiddleware())
	authAPI.RegisterPrivateRoutes(private)

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
