package echoEntrypoint

import (
	hasher "github.com/dylanbatar/simple-login-system/internal/infra/adapters/hasher"
	"github.com/dylanbatar/simple-login-system/internal/infra/adapters/jwt"
	"github.com/dylanbatar/simple-login-system/internal/infra/entrypoints/apiRest/echoEntrypoint/handlers"
	"github.com/dylanbatar/simple-login-system/internal/infra/factory/repository"
	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"log"
)

type EchoApiEntryPoint struct {
	echoInstance *echo.Echo
	handler      handlers.Handler
}

func NewEchoApiRest() *EchoApiEntryPoint {
	e := echo.New()
	validatorInstance := validator.New()
	hasherInstance := hasher.NewHasherAdapter()
	repositoryInstance, _ := repository.NewRepositoryFactory("mongo")
	jwtInstance := jwt.NewJwtAdapter()
	handlerInstace := handlers.NewHandler(validatorInstance, hasherInstance, repositoryInstance, jwtInstance)
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
