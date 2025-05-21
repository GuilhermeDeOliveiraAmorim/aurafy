package repositoriesinterfaces

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type SpotifyUser struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Country     string `json:"country"`
	Product     string `json:"product"`
	Images      []struct {
		URL string `json:"url"`
	} `json:"images"`
}

type SpotifyRepositoryInterface interface {
	ExchangeCodeForToken(code string) (TokenResponse, error)
	GetSpotifyUserByAccessToken(accessToken string) (SpotifyUser, error)
}
