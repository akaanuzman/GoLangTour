// controllers/user.go
package controllers

import (
	"net/http"
	"note-app/src/models"
	"note-app/src/models/requests"
	"note-app/src/services"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
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

func (controller *AuthController) ChangePassword(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)

	userIDHex, ok := claims["user_id"].(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid user ID"})
	}

	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid user ID"})
	}

	var req requests.PasswordChangeRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := controller.authService.ChangePassword(userID, &req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "password changed successfully",
	})
}

func (controller *AuthController) ForgotPassword(c echo.Context) error {

	var req requests.ForgotPasswordRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := controller.authService.ForgotPassword(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "reset password link sent to your email",
	})
}
