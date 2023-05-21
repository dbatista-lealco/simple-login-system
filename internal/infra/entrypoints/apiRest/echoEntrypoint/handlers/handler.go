package handlers

import "github.com/go-playground/validator/v10"

type Handler struct {
	validator *validator.Validate
}

type ResponseMessage struct {
	Message string `json:"message"`
}

func NewHandler() *Handler {
	return &Handler{
		validator: validator.New(),
	}
}
