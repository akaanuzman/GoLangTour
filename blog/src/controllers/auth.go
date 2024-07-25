// controllers/user.go
package controllers

import (
	"blog/src/models"
	"blog/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authService services.IAuthService
}

func NewUserController(authService services.IAuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (controller *AuthController) Register(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := controller.authService.Register(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (controller *AuthController) Login(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	token, dbUser, err := controller.authService.Login(&user)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
		"user":  dbUser,
	})
}
