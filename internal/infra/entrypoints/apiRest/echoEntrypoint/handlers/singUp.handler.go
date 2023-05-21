package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (handler *Handler) SignUp(c echo.Context) error {
	return c.String(http.StatusOK, "Registered")
}
