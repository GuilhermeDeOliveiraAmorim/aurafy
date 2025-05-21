package handlers

import (
	"github.com/GuilhermeDeOliveiraAmorim/aurafy/internal/database"
	"github.com/GuilhermeDeOliveiraAmorim/aurafy/internal/factories"
)

type HandlerFactory struct {
	SpotifyHandler *SpotifyHandler
}

func NewHandlerFactory(inputFactory database.StorageInput) *HandlerFactory {
	spotifyFactory := factories.NewSpotifyFactory(inputFactory)

	return &HandlerFactory{
		SpotifyHandler: NewSpotifyHandler(spotifyFactory),
	}
}
