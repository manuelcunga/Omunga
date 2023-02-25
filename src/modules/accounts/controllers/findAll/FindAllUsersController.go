package controller

import (
	"net/http"

	"github.com/labstack/echo"
	usecases "github.com/manuelcunga/Omunga/src/modules/accounts/usecases/findAll"
)

type UsersController struct {
	findAllUsersUseCases usecases.FindAllUsersUseCases
}

func NewUsersController(usersUcaseCase usecases.FindAllUsersUseCases) UsersController {
	return UsersController{
		findAllUsersUseCases: usersUcaseCase,
	}
}

func (userController *UsersController) FindAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		userList, err := userController.findAllUsersUseCases.FindAll()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"Erro ao listar todos usuários.": err.Error()})
		}

		return c.JSON(http.StatusOK, userList)
	}
}
