package controller

import (
	"net/http"

	"github.com/labstack/echo"
	usecases "github.com/manuelcunga/Omunga/src/modules/accounts/usecases/findOne"
)

type FindOneUserController struct {
	findAllUsersUseCases usecases.FindOndeUserUseCases
}

func NewFindOneUserController(usersUcaseCase usecases.FindOndeUserUseCases) FindOneUserController {
	return FindOneUserController{
		findAllUsersUseCases: usersUcaseCase,
	}
}

func (userController *FindOneUserController) FindOneUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("id")
		user, err := userController.findAllUsersUseCases.FindOne(userId)

		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, user)
	}
}
