package dto

type PasswordResetDTO struct {
	Email            string `json:"email" validate:"required,email"`
	VerificationCode string `json:"verificationCode" validate:"required"`
	NewPassword      string `json:"newPassword" validate:"required""`
}
