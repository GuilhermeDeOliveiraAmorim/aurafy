package handlers

import (
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/aurafy/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/aurafy/internal/factories"
	"github.com/GuilhermeDeOliveiraAmorim/aurafy/internal/usecases"
	"github.com/gin-gonic/gin"
)

type SpotifyHandler struct {
	SpotifyFactory *factories.SpotifyFactory
}

func NewSpotifyHandler(factory *factories.SpotifyFactory) *SpotifyHandler {
	return &SpotifyHandler{
		SpotifyFactory: factory,
	}
}

// @Summary Create a new championship
// @Description Registers a new championship in the system
// @Tags Championship
// @Accept json
// @Produce json
// @Param request body usecases.CreateUserInputDTO true "Championship data"
// @Success 201 {object} usecases.CreateUserOutputDTO
// @Failure 400 {object} exceptions.ProblemDetails "Bad Request"
// @Failure 500 {object} exceptions.ProblemDetails "Internal Server Error"
// @Failure 401 {object} exceptions.ProblemDetails "Unauthorized"
// @Security BearerAuth
// @Router /oauth [post]
func (h *SpotifyHandler) OAuth(c *gin.Context) {
	var oAuth usecases.CreateUserInputDTO
	if err := c.ShouldBindJSON(&oAuth); err != nil {
		problem := exceptions.NewProblemDetails(exceptions.InternalServerError, exceptions.ErrorMessage{
			Title:  "Erro ao fazer o parser",
			Detail: "Erro ao fazer o parser do campeonato",
		})

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": problem})
		return
	}

	input := usecases.CreateUserInputDTO{
		Code: oAuth.Code,
	}

	output, errs := h.SpotifyFactory.OAuth.Execute(input)
	if len(errs) > 0 {
		exceptions.HandleErrors(c, errs)
		return
	}

	c.JSON(http.StatusCreated, output)
}
