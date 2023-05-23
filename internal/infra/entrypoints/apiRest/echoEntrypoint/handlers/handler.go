package handlers

import (
	"github.com/dylanbatar/simple-login-system/internal/application/ports"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	validator  *validator.Validate
	hasher     ports.IHasher
	repository ports.IRepository
}

type ResponseMessage struct {
	Message string `json:"message"`
}

func NewHandler(validator *validator.Validate, hasher ports.IHasher, repository ports.IRepository) *Handler {
	return &Handler{
		validator:  validator,
		hasher:     hasher,
		repository: repository,
	}
}
