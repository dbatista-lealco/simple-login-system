package echoEntrypoint

import (
	"github.com/dylanbatar/simple-login-system/internal/infra/entrypoints/apiRest/echoEntrypoint/handlers"
	"github.com/labstack/echo/v4"
	"log"
)

type EchoApiEntryPoint struct {
	echoInstance *echo.Echo
	handler      handlers.Handler
}

func NewEchoApiRest() *EchoApiEntryPoint {
	e := echo.New()
	handlerInstace := handlers.NewHandler()
	return &EchoApiEntryPoint{
		echoInstance: e,
		handler:      *handlerInstace,
	}
}

func (api *EchoApiEntryPoint) RunServer() {
	api.setupRoutes()
	log.Fatalln(api.echoInstance.Start(":3000"))
}

func (api *EchoApiEntryPoint) setupRoutes() {
	api.echoInstance.GET("/health", api.handler.Health)

	api.echoInstance.POST("/signIn", api.handler.SignIn)
	api.echoInstance.POST("/signUp", api.handler.SignUp)
	api.echoInstance.POST("/signUp", api.handler.SignUp)
	api.echoInstance.POST("/password/reset", api.handler.ResetPassword)
}
