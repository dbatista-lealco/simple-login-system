package handlers

import (
	"fmt"
	"github.com/dylanbatar/simple-login-system/internal/infra/entrypoints/dto"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (handler *Handler) SignUp(c echo.Context) error {
	user := dto.SignUpDTO{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseMessage{Message: err.Error()})
	}

	if err := handler.validator.Struct(user); err != nil {
		// TODO: Create a util to send a clients witch field's are failing
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println("- Field:", err.Field())
			fmt.Println("- Error:", err.Tag())
			fmt.Println()
		}

		return c.JSON(http.StatusBadRequest, ResponseMessage{Message: "Invalid request"})
	}

	return c.JSON(http.StatusOK, ResponseMessage{Message: "User signup successfully"})
}
