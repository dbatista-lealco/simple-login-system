package api

import (
	"fmt"
	"github.com/dylanbatar/simple-login-system/internal/infra/entrypoints/apiRest/echoEntrypoint"
	"github.com/dylanbatar/simple-login-system/internal/infra/factory/api/enum"
	"github.com/dylanbatar/simple-login-system/internal/infra/factory/ports"
)

func NewApiFactory(provider enum.ApiRestProvider) (ports.IApiFactory, error) {
	switch provider {
	case enum.EchoProvider:
		return echoEntrypoint.NewEchoApiRest(), nil
	}
	return nil, fmt.Errorf("provider not found")
}
