package routes

import (
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/aurafy/internal/config"
	"github.com/GuilhermeDeOliveiraAmorim/aurafy/internal/database"
	"github.com/GuilhermeDeOliveiraAmorim/aurafy/internal/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(storageInput database.StorageInput) *gin.Engine {
	handlerFactory := handlers.NewHandlerFactory(storageInput)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.FRONT_END_URL_VAR.FRONT_END_URL_DEV, config.FRONT_END_URL_VAR.FRONT_END_URL_PROD},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	public := r.Group("/")
	{
		public.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		public.POST("oauth", handlerFactory.SpotifyHandler.OAuth)
	}

	return r
}
