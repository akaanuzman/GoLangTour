package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// JWTMiddleware is a middleware that checks if the request has a valid jwt.
func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
		Claims:     jwt.MapClaims{},
		ErrorHandlerWithContext: func(err error, c echo.Context) error {
			c.Logger().Error("JWT Error:", err)
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid or expired jwt"})
		},
	})
}
