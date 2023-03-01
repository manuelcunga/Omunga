package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	usecases "github.com/manuelcunga/Omunga/src/modules/accounts/usecases/me"
)

type MeController struct {
	meUseCases usecases.MeUseCases
}

func NewMeController(meUseCases usecases.MeUseCases) MeController {
	return MeController{meUseCases: meUseCases}
}

func (meController *MeController) Me() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, ok := c.Get("id").(string)
		if !ok {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "userID not found or is not a string"})
		}

		me, err := meController.meUseCases.Me(userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, me)
	}
}
