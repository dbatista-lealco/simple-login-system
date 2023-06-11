package ports

import "github.com/dylanbatar/simple-login-system/internal/domain"

type IRepository interface {
	Save(user domain.User) error
	FindByEmail(email string) (domain.User, error)
	VerifyResetPassword(verificationCode, email string) (domain.PasswordReset, error)
	UpdatePassword(email, newPassword string) error
	GenerateOtp(data domain.PasswordReset) error
	InactiveOtp(email string) error
}
