package generateOTP

import (
	"fmt"
	"github.com/dylanbatar/simple-login-system/internal/application/ports"
	"github.com/dylanbatar/simple-login-system/internal/domain"
	"github.com/dylanbatar/simple-login-system/internal/domain/utils"
)

type GenerateOTPUseCase struct {
	Repository ports.IRepository
}

func NewGeneateOTPUseCase(repository ports.IRepository) *GenerateOTPUseCase {
	return &GenerateOTPUseCase{
		Repository: repository,
	}
}

func (u *GenerateOTPUseCase) GenerateOTP(email string) (domain.PasswordReset, error) {
	user, err := u.Repository.FindByEmail(email)

	if user.Email == "" {
		return domain.PasswordReset{}, fmt.Errorf("account not found %s", err.Error())
	}

	if err != nil {
		return domain.PasswordReset{}, err
	}

	err = u.Repository.InactiveOtp(email)

	if err != nil {
		return domain.PasswordReset{}, fmt.Errorf("error inactivating otp code")
	}

	otp := utils.GenerateVerifyCode(4)

	entry := domain.NewPasswordReset(otp, email)

	err = u.Repository.GenerateOtp(entry)

	if err != nil {
		return domain.PasswordReset{}, fmt.Errorf("error generating otp code %s", err.Error())
	}

	return entry, nil
}
