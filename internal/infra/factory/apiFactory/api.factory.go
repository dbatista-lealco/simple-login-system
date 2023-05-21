package apiFactory

import (
	"fmt"
	"github.com/dylanbatar/simple-login-system/internal/infra/entrypoints/apiRest"
	"github.com/dylanbatar/simple-login-system/internal/infra/factory/apiFactory/enum"
	"github.com/dylanbatar/simple-login-system/internal/infra/factory/ports"
)

func NewApiFactory(provider enum.ApiRestProvider) (ports.IApiFactory, error) {
	switch provider {
	case enum.EchoProvider:
		return apiRest.NewEchoApiRest(), nil
	}
	return nil, fmt.Errorf("provider not found")
}
