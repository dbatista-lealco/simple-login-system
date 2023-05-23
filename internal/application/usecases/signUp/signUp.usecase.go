package signUp

import (
	"github.com/dylanbatar/simple-login-system/internal/application/ports"
	"github.com/dylanbatar/simple-login-system/internal/domain"
	"github.com/dylanbatar/simple-login-system/internal/dto"
)

type SignUpUsecase struct {
	repository ports.IRepository
	hasher     ports.IHasher
}

func NewSignUpUsecase(repository ports.IRepository, hasher ports.IHasher) *SignUpUsecase {
	return &SignUpUsecase{
		repository: repository,
		hasher:     hasher,
	}
}

func (u *SignUpUsecase) SignUp(newUser dto.SignUpDTO) error {
	hashPassword, err := u.hasher.Hash(newUser.Password)

	if err != nil {
		return err
	}

	user, err := domain.NewUser(newUser.Name, newUser.Email, hashPassword)

	if err != nil {
		return err
	}

	return u.repository.Save(*user)
}
