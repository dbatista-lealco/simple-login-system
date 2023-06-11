package handlers

import (
	"errors"
	"github.com/dylanbatar/simple-login-system/internal/application/usecases/resetPassword"
	"github.com/dylanbatar/simple-login-system/internal/dto"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (handler *Handler) ResetPassword(c echo.Context) error {
	var req dto.PasswordResetDTO

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseMessage{Message: err.Error()})
	}

	if err := handler.validator.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseMessage{Message: "Invalid params"})
	}

	userCase := resetPassword.NewPasswordResetUseCase(handler.hasher, handler.repository)

	message, err := userCase.Reset(req.VerificationCode, req.Email, req.NewPassword)

	if err != nil {
		switch {
		case errors.Is(err, resetPassword.VerificationCodeError):
			return c.JSON(http.StatusBadRequest, ResponseMessage{Message: resetPassword.VerificationCodeError.Error()})
		default:
			return c.JSON(http.StatusInternalServerError, ResponseMessage{Message: err.Error()})
		}
	}

	return c.JSON(http.StatusOK, ResponseMessage{Message: message})
}
