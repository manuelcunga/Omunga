package routes

import (
	"github.com/labstack/echo"
	controller "github.com/manuelcunga/Omunga/src/modules/accounts/controllers/create"
	"github.com/manuelcunga/Omunga/src/modules/accounts/factory"
)

func UserRoutes(e *echo.Group) {
	userFactory := factory.NewUserUseCase()
	createUserController := controller.NewUserController(userFactory)
	e.POST("/user/create", createUserController.CreateUser())
}
