package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

type WelcomeResponse struct {
	Message string `json:"message"`
}

func Home(app *echo.Echo) {
	app.GET("/api/v1", func(c echo.Context) error {
		response := &WelcomeResponse{Message: "Bem-vindo ao Omunga Back-Ends!"}
		return c.JSON(http.StatusOK, response)
	})
}

func Setup(app *echo.Echo) {
	UserRoutes(app)
}
