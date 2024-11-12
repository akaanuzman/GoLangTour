// controllers/user.go
package controllers

import (
	"chat-app/src/models"
	"chat-app/src/models/requests"
	"chat-app/src/services"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AuthController is a struct that defines the methods of the auth controller.
type AuthController struct {
	authService services.IAuthService
}

// NewUserController creates a new instance of AuthController with the provided IAuthService.
func NewUserAuthController(authService services.IAuthService) *AuthController {
	return &AuthController{authService: authService}
}

// Register registers a new user.
func (controller *AuthController) Register(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user.CreatedAt = primitive.NewDateTimeFromTime(time.Now())

	if err := controller.authService.Register(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

// Login logs in a user.
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

// ChangePassword changes the password of a user.
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

// ForgotPassword sends a reset password link to the user's email.
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
