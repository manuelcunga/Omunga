package routes

import (
	"github.com/labstack/echo"
	controller "github.com/manuelcunga/Omunga/src/modules/accounts/controllers/create"
	usersController "github.com/manuelcunga/Omunga/src/modules/accounts/controllers/findAll"
	findOneUserController "github.com/manuelcunga/Omunga/src/modules/accounts/controllers/findOne"
	updateUserController "github.com/manuelcunga/Omunga/src/modules/accounts/controllers/update"
	"github.com/manuelcunga/Omunga/src/modules/accounts/factory"
)

func UserRoutes(e *echo.Group) {
	userFactory := factory.NewUserUseCase()
	createUserController := controller.NewUserController(userFactory)

	listUsersFactory := factory.NewListUsersUseCase()
	listUsersUseCaseController := usersController.NewUsersController(listUsersFactory)

	findOneuserFactory := factory.NewFindOneUserUseCases()
	findOneUserController := findOneUserController.NewFindOneUserController(findOneuserFactory)

	udateUserFactory := factory.NewUpdateUserUsecases()
	updateUser := updateUserController.NewUpdateUserController(udateUserFactory)

	e.POST("/user/create", createUserController.CreateUser())
	e.GET("/users", listUsersUseCaseController.FindAllUsers())
	e.GET("/user/:id", findOneUserController.FindOneUser())
	e.PATCH("/user/:id", updateUser.UpdateUser())
}
