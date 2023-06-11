package dto

type GenerateOtpDTO struct {
	Email string `json:"email" validate:"required,email"`
}
