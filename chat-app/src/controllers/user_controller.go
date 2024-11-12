package controllers

import (
	"chat-app/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UserController is a struct that defines the methods of the user controller.
type UserController struct {
	userService services.IUserService
}

// NewUserController creates a new instance of UserController with the provided IUserService.
func NewUserController(userService services.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// GetUserByID gets a user by ID.
func (controller *UserController) GetUserByID(c echo.Context) error {
	userID := c.Param("id")
	user, err := controller.userService.GetUserByID(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
	}
	return c.JSON(http.StatusOK, user)
}

// GetUserByEmail gets a user by email.
func (controller *UserController) GetUserByEmail(c echo.Context) error {
	email := c.QueryParam("email")
	user, err := controller.userService.GetUserByEmail(email)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
	}
	return c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user.
func (controller *UserController) DeleteUser(c echo.Context) error {
	userID := c.Param("id")
	if err := controller.userService.DeleteUser(userID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete user"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "User deleted successfully"})
}
