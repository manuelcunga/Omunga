package routes

import (
	"github.com/labstack/echo"
	controller "github.com/manuelcunga/Omunga/src/modules/accounts/controllers/create"
	usersController "github.com/manuelcunga/Omunga/src/modules/accounts/controllers/findAll"
	"github.com/manuelcunga/Omunga/src/modules/accounts/factory"
)

func UserRoutes(e *echo.Group) {
	userFactory := factory.NewUserUseCase()
	createUserController := controller.NewUserController(userFactory)

	listUsersFactory := factory.NewListUsersUseCase()
	listUsersUseCaseController := usersController.NewUsersController(listUsersFactory)

	e.POST("/user/create", createUserController.CreateUser())
	e.GET("/users", listUsersUseCaseController.FindAllUsers())
}
