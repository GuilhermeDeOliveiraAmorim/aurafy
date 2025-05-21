package repositoriesimpl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/GuilhermeDeOliveiraAmorim/aurafy/internal/config"
	repositoriesinterfaces "github.com/GuilhermeDeOliveiraAmorim/aurafy/internal/repositories_interfaces"
	"gorm.io/gorm"
)

type SpotifyRepository struct {
	gorm *gorm.DB
}

func NewSpotifyRepository(gorm *gorm.DB) *SpotifyRepository {
	return &SpotifyRepository{
		gorm: gorm,
	}
}

func (r *SpotifyRepository) ExchangeCodeForToken(code string) (repositoriesinterfaces.TokenResponse, error) {
	endpoint := "https://accounts.spotify.com/api/token"

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", config.SPOTIFY_AUTH_SERVICE.REDIRECT_URI)
	data.Set("client_id", config.SPOTIFY_AUTH_SERVICE.CLIENT_ID)
	data.Set("client_secret", config.SPOTIFY_AUTH_SERVICE.CLIENT_SECRET)

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return repositoriesinterfaces.TokenResponse{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return repositoriesinterfaces.TokenResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return repositoriesinterfaces.TokenResponse{}, errors.New("failed to get token from Spotify")
	}

	var tokenResp repositoriesinterfaces.TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return repositoriesinterfaces.TokenResponse{}, err
	}

	return tokenResp, nil
}

func (r *SpotifyRepository) GetSpotifyUserByAccessToken(accessToken string) (repositoriesinterfaces.SpotifyUser, error) {
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
	if err != nil {
		return repositoriesinterfaces.SpotifyUser{}, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return repositoriesinterfaces.SpotifyUser{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return repositoriesinterfaces.SpotifyUser{}, fmt.Errorf("erro ao buscar user: %s", string(bodyBytes))
	}

	var user repositoriesinterfaces.SpotifyUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return repositoriesinterfaces.SpotifyUser{}, err
	}

	return user, nil
}
