package usecases

import (
	"fmt"

	"github.com/GuilhermeDeOliveiraAmorim/aurafy/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/aurafy/internal/exceptions"
	repositoriesinterfaces "github.com/GuilhermeDeOliveiraAmorim/aurafy/internal/repositories_interfaces"
)

type CreateUserInputDTO struct {
	Code string `json:"code"`
}

type CreateUserOutputDTO struct {
	SuccessMessage string `json:"success_message"`
	ContentMessage string `json:"content_message"`
}

type CreateUser struct {
	SpotifyRepository repositoriesinterfaces.SpotifyRepositoryInterface
	ImageRepository   repositoriesinterfaces.ImageRepositoryInterface
}

func NewCreateUserUseCase(
	SpotifyRepository repositoriesinterfaces.SpotifyRepositoryInterface,
	ImageRepository repositoriesinterfaces.ImageRepositoryInterface,
) *CreateUser {
	return &CreateUser{
		SpotifyRepository: SpotifyRepository,
		ImageRepository:   ImageRepository,
	}
}

func (u *CreateUser) Execute(input CreateUserInputDTO) (CreateUserOutputDTO, []exceptions.ProblemDetails) {
	var problems []exceptions.ProblemDetails

	tokenResp, err := u.SpotifyRepository.ExchangeCodeForToken(input.Code)
	if err != nil {
		problem := exceptions.NewProblemDetails(exceptions.InternalServerError, exceptions.ErrorMessage{
			Title:  "Erro ao obter token",
			Detail: err.Error(),
		})

		problems = append(problems, problem)

		return CreateUserOutputDTO{}, problems
	}

	user, err := u.SpotifyRepository.GetSpotifyUserByAccessToken(tokenResp.AccessToken)
	if err != nil {
		problem := exceptions.NewProblemDetails(exceptions.InternalServerError, exceptions.ErrorMessage{
			Title:  "Erro ao usuário spotify",
			Detail: err.Error(),
		})

		problems = append(problems, problem)

		return CreateUserOutputDTO{}, problems
	}

	profileImage, saveImageErr := u.ImageRepository.SaveImage(user.Images[0].URL)
	if saveImageErr != nil {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.InternalServerError, exceptions.ErrorMessage{
			Title:  "Erro ao criar imagem do usuário",
			Detail: saveImageErr.Error(),
		}))

		return CreateUserOutputDTO{}, problems
	}

	newUser := entities.NewUser(entities.CreateUser{
		SpotifyID:    user.ID,
		DisplayName:  user.DisplayName,
		Email:        user.Email,
		Country:      user.Country,
		Product:      user.Product,
		ProfileImage: profileImage,
		AccessToken:  tokenResp.AccessToken,
		RefreshToken: tokenResp.RefreshToken,
	})

	fmt.Println("newUser", newUser)

	return CreateUserOutputDTO{
		SuccessMessage: "Usuário autenticado com sucesso!",
		ContentMessage: "Token salvo e usuário registrado.",
	}, problems
}
