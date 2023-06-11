package handlers

import (
	"github.com/dylanbatar/simple-login-system/internal/application/usecases/generateOTP"
	"github.com/dylanbatar/simple-login-system/internal/dto"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (handler *Handler) GenerateOtp(c echo.Context) error {
	var req dto.GenerateOtpDTO

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseMessage{Message: err.Error()})
	}

	if err := handler.validator.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseMessage{Message: "Invalid params"})
	}

	useCase := generateOTP.NewGeneateOTPUseCase(handler.repository)

	result, err := useCase.GenerateOTP(req.Email)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   result.Code,
		"active": result.Active,
		"email":  result.Email,
	})
}
