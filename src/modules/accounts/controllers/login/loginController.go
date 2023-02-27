package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/manuelcunga/Omunga/src/modules/accounts/dtos"
	usecases "github.com/manuelcunga/Omunga/src/modules/accounts/usecases/login"
)

type LoginController struct {
	loginUseCases usecases.LoginUseCases
}

func NewLoginController(loginUseCases usecases.LoginUseCases) *LoginController {
	return &LoginController{loginUseCases: loginUseCases}
}

func (loginController *LoginController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input dtos.LoginDTO
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		}

		output, err := loginController.loginUseCases.Login(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusCreated, output)
	}
}
