package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (handler *Handler) ResetPassword(c echo.Context) error {
	return c.String(http.StatusOK, "reset")
}
