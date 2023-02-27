package controller

import (
	"net/http"

	"github.com/labstack/echo"
	usecases "github.com/manuelcunga/Omunga/src/modules/accounts/usecases/delete"
)

type DeleteUserController struct {
	deleteUserController usecases.DeleteUserUseCases
}

func NewDeleteUserController(usersUcaseCase usecases.DeleteUserUseCases) DeleteUserController {
	return DeleteUserController{
		deleteUserController: usersUcaseCase,
	}
}

func (userController *DeleteUserController) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("id")

		err := userController.deleteUserController.Execute(userId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Usuário excluído com sucesso"})
	}
}
