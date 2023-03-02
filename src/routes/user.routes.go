package routes

import (
	"github.com/labstack/echo"
	controller "github.com/manuelcunga/Omunga/src/modules/accounts/controllers/create"
	deleteUserController "github.com/manuelcunga/Omunga/src/modules/accounts/controllers/delete"
	usersController "github.com/manuelcunga/Omunga/src/modules/accounts/controllers/findAll"
	findOneUserController "github.com/manuelcunga/Omunga/src/modules/accounts/controllers/findOne"
	loginController "github.com/manuelcunga/Omunga/src/modules/accounts/controllers/login"
	meController "github.com/manuelcunga/Omunga/src/modules/accounts/controllers/me"
	userProfileController "github.com/manuelcunga/Omunga/src/modules/accounts/controllers/profile"
	updateUserController "github.com/manuelcunga/Omunga/src/modules/accounts/controllers/update"
	"github.com/manuelcunga/Omunga/src/modules/accounts/factory"
	"github.com/manuelcunga/Omunga/src/modules/accounts/middlewares"
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

	deleteUserFactory := factory.NewDeleteUserUseCases()
	deleteUser := deleteUserController.NewDeleteUserController(deleteUserFactory)

	loginFactory := factory.NewLoginUseCases()
	loginUser := loginController.NewLoginController(loginFactory)

	profileFactory := factory.NewProfileUseCases()
	profile := userProfileController.NewProfileController(profileFactory)

	meFactory := factory.NewMeUseCases()
	meUserController := meController.NewMeController(meFactory)

	e.POST("/auth/login", loginUser.Login())
	e.POST("/user/create", createUserController.CreateUser())

	e.Use(middlewares.AuthMiddlewareFunc())
	e.GET("/users", listUsersUseCaseController.FindAllUsers())
	e.GET("/user/:id", findOneUserController.FindOneUser())
	e.PATCH("/user/:id", updateUser.UpdateUser())
	e.DELETE("/user/:id", deleteUser.DeleteUser())
	e.GET("/user/profile/:id", profile.Profile())
	e.GET("/user/me", meUserController.Me())

}
