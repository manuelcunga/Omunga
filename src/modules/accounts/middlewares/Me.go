package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Me(c echo.Context) error {
	user := c.Get("user")
	if user != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{"user": user})
	} else {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}
}
