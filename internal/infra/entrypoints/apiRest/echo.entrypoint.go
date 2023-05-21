package apiRest

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type EchoApiEntryPoint struct {
}

func NewEchoApiRest() *EchoApiEntryPoint {
	return &EchoApiEntryPoint{}
}

func (api *EchoApiEntryPoint) RunServer() {
	e := echo.New()
	api.setupRoutes(e)
	log.Fatalln(e.Start(":3000"))
}

func (api *EchoApiEntryPoint) setupRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hola cv")
	})
}
