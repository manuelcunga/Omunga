package controller

import (
	"net/http"

	"github.com/labstack/echo"
	usecases "github.com/manuelcunga/Omunga/src/modules/accounts/usecases/profile"
)

type ProfileUserController struct {
	profileUserUsecases usecases.ProfileUseCases
}

func NewProfileController(profileUserCase usecases.ProfileUseCases) ProfileUserController {
	return ProfileUserController{
		profileUserUsecases: profileUserCase,
	}
}

func (profileController *ProfileUserController) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("id")

		profile, err := profileController.profileUserUsecases.Profile(userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, profile)
	}
}
