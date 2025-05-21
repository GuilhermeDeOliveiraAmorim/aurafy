package factories

import (
	"github.com/GuilhermeDeOliveiraAmorim/aurafy/internal/database"
	repositoriesimpl "github.com/GuilhermeDeOliveiraAmorim/aurafy/internal/repositories_impl"
	"github.com/GuilhermeDeOliveiraAmorim/aurafy/internal/usecases"
)

type SpotifyFactory struct {
	OAuth *usecases.CreateUser
}

func NewSpotifyFactory(input database.StorageInput) *SpotifyFactory {
	spotifyRepository := repositoriesimpl.NewSpotifyRepository(input.DB)
	imageRepository := repositoriesimpl.NewImageRepository(input.BucketName)

	oAuth := usecases.NewCreateUserUseCase(spotifyRepository, imageRepository)

	return &SpotifyFactory{
		OAuth: oAuth,
	}
}
