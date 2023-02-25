package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

type WelcomeResponse struct {
	Message string `json:"message"`
}

func Home(app *echo.Group) {
	app.GET("/", func(c echo.Context) error {
		response := &WelcomeResponse{
			Message: "Bem-vindo ao Omunga Back-Ends!"}
		return c.JSON(http.StatusOK, response)
	})
}

func Setup(app *echo.Echo) {
	api := app.Group("/api/v1")
	Home(api)
	UserRoutes(api)
}
