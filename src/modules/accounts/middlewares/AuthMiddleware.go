package middlewares

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.String(http.StatusUnauthorized, "É necessário estar logado, faça o login")
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			return c.String(http.StatusUnauthorized, "Token inválido")
		}

		if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
			userID, err := strconv.Atoi(claims.Subject)
			if err != nil {
				return c.String(http.StatusInternalServerError, "ID de usuário inválido no token")
			}

			if err != nil {
				return c.String(http.StatusInternalServerError, "ID de usuário inválido no token")
			}

			ctx := context.WithValue(c.Request().Context(), "userID", userID)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		} else {
			return c.String(http.StatusUnauthorized, "Token inválido")
		}
	}
}

func AuthMiddlewareFunc() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return AuthMiddleware(next)(c)
		}
	}
}
