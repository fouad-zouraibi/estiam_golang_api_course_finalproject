package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/phramos07/finalproject/services"
	"github.com/phramos07/finalproject/types"
)

type UserHandler interface {
	Create(echo.Context) error
}

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) Create(c echo.Context) error {
	user := new(types.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	// Call the service to create a new user
	if err := h.userService.CreateNewUser(c.Request().Context(), user); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}