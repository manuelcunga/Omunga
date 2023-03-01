package middlewares

import (
	"net/http"

	"github.com/labstack/echo"
)

func MeMiddlewareFunc() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c := r.Context()
			user := c.Value("user")
			if user != nil {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}
		})
	}
}

func Me() echo.MiddlewareFunc {
	return echo.WrapMiddleware(MeMiddlewareFunc())
}
