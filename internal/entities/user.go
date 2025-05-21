package entities

type User struct {
	SharedEntity
	SpotifyID    string
	DisplayName  string
	Email        string
	Country      string
	Product      string
	ProfileImage string
	AccessToken  string
	RefreshToken string
}

type CreateUser struct {
	SpotifyID    string
	DisplayName  string
	Email        string
	Country      string
	Product      string
	ProfileImage string
	AccessToken  string
	RefreshToken string
}

func NewUser(input CreateUser) *User {
	return &User{
		SharedEntity: *NewSharedEntity(),
		SpotifyID:    input.SpotifyID,
		DisplayName:  input.DisplayName,
		Email:        input.Email,
		Country:      input.Country,
		Product:      input.Product,
		ProfileImage: input.ProfileImage,
		AccessToken:  input.AccessToken,
		RefreshToken: input.RefreshToken,
	}
}
