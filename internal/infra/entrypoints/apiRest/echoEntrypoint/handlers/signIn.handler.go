package handlers

import (
	"errors"
	"github.com/dylanbatar/simple-login-system/internal/application/usecases/signIn"
	"github.com/dylanbatar/simple-login-system/internal/dto"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (handler *Handler) SignIn(c echo.Context) error {
	userDto := dto.SignInDTO{}

	if err := c.Bind(&userDto); err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseMessage{Message: err.Error()})
	}

	if err := handler.validator.Struct(userDto); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseMessage{Message: "Invalid params"})
	}

	useCase := signIn.NewSignInUseCase(handler.repository, handler.hasher, handler.JWT)

	token, err := useCase.SignIn(userDto)

	if err != nil {
		switch {
		case errors.Is(err, signIn.AccountNotFound), errors.Is(err, signIn.InCorrectCredentials):
			return c.JSON(http.StatusBadRequest, ResponseMessage{Message: err.Error()})
		default:
			return c.JSON(http.StatusInternalServerError, ResponseMessage{Message: err.Error()})
		}
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "Login Successfully",
		"token":   string(token),
	})

	return nil
}
