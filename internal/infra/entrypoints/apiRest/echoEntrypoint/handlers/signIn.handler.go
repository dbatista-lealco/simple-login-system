package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (handler *Handler) SignIn(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome")
}
