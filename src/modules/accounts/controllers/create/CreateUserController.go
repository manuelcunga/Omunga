package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/manuelcunga/Omunga/src/modules/accounts/dtos"
	usecases "github.com/manuelcunga/Omunga/src/modules/accounts/usecases/create"
)

type UsersController struct {
	userUseCases usecases.CreateUserUseCases
}

func NewUserController(userUcaseCase usecases.CreateUserUseCases) UsersController {
	return UsersController{
		userUseCases: userUcaseCase,
	}
}

func (userController *UsersController) CreateUser() echo.HandlerFunc {

	return func(c echo.Context) error {
		var input dtos.CreateUserDTO
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		}

		output, err := userController.userUseCases.Execute(&input)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusCreated, output)
	}
}
