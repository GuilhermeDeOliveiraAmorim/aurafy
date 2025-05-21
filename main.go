package main

import (
	"context"

	"github.com/GuilhermeDeOliveiraAmorim/aurafy/internal/config"
	"github.com/GuilhermeDeOliveiraAmorim/aurafy/internal/database"
	"github.com/GuilhermeDeOliveiraAmorim/aurafy/internal/models"
	"github.com/GuilhermeDeOliveiraAmorim/aurafy/internal/routes"

	_ "github.com/GuilhermeDeOliveiraAmorim/aurafy/api"
)

func main() {
	ctx := context.Background()

	db, sqlDB := database.SetupDatabaseConnection(ctx, database.LOCAL)

	models.Migration(db, sqlDB)

	router := routes.SetupRouter(database.StorageInput{
		DB:         db,
		BucketName: config.GOOGLE_VAR.IMAGE_BUCKET_NAME,
	})

	router.Run(":8080")
}
