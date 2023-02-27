package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/manuelcunga/Omunga/src/modules/accounts/dtos"
	usecases "github.com/manuelcunga/Omunga/src/modules/accounts/usecases/update"
)

type UpdateUserController struct {
	updateUserController usecases.UpdateUserUseCases
}

func NewUpdateUserController(usersUcaseCase usecases.UpdateUserUseCases) UpdateUserController {
	return UpdateUserController{
		updateUserController: usersUcaseCase,
	}
}

func (userController *UpdateUserController) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("id")

		var userInput dtos.UpdateUserDTO
		err := c.Bind(&userInput)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Entrada de dados inválida"})
		}

		user, err := userController.updateUserController.Execute(userId, &userInput)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, user)
	}
}
