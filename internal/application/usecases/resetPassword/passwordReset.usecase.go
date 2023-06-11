package resetPassword

import (
	"errors"
	"fmt"
	"github.com/dylanbatar/simple-login-system/internal/application/ports"
)

var VerificationCodeError error = errors.New("invalid verification code")
var ChangePasswordError error = errors.New("change password wasn't possible")

type PasswordResetUseCase struct {
	Hasher     ports.IHasher
	Repository ports.IRepository
}

func NewPasswordResetUseCase(hasher ports.IHasher, repository ports.IRepository) *PasswordResetUseCase {
	return &PasswordResetUseCase{
		Hasher:     hasher,
		Repository: repository,
	}
}

func (u *PasswordResetUseCase) Reset(verificationCode, email, newPassword string) (string, error) {
	resetPasswordRequest, err := u.Repository.VerifyResetPassword(verificationCode, email)

	if err != nil {
		return "", fmt.Errorf("error finding verification code given %s", err.Error())
	}

	if resetPasswordRequest.Code == "" || !resetPasswordRequest.Active {
		return "", VerificationCodeError
	}

	hashPassword, _ := u.Hasher.Hash(newPassword)

	err = u.Repository.UpdatePassword(email, hashPassword)

	if err != nil {
		return "", ChangePasswordError
	}

	return "password changed successfully", nil
}
