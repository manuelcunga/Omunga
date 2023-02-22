package middlewares

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.String(http.StatusUnauthorized, "Authorization header is required")
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			return c.String(http.StatusUnauthorized, "Invalid token")
		}

		if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
			userID, err := strconv.Atoi(claims.Subject)
			if err != nil {
				return c.String(http.StatusInternalServerError, "Invalid user ID in token")
			}

			ctx := context.WithValue(c.Request().Context(), "userID", userID)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		} else {
			return c.String(http.StatusUnauthorized, "Invalid token")
		}
	}
}
