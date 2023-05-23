package dto

type SignUpDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required,min=6"`
	Password string `json:"password" validate:"required"`
}
