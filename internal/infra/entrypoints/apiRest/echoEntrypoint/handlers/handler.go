package handlers

import (
	"github.com/dylanbatar/simple-login-system/internal/application/ports"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	validator  *validator.Validate
	hasher     ports.IHasher
	repository ports.IRepository
	JWT        ports.IJWT
}

type ResponseMessage struct {
	Message string `json:"message"`
}

func NewHandler(validator *validator.Validate, hasher ports.IHasher, repository ports.IRepository, jwt ports.IJWT) *Handler {
	return &Handler{
		validator:  validator,
		hasher:     hasher,
		repository: repository,
		JWT:        jwt,
	}
}
