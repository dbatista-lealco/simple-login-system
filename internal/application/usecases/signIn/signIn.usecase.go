package signIn

import (
	"errors"
	"github.com/dylanbatar/simple-login-system/internal/application/ports"
	"github.com/dylanbatar/simple-login-system/internal/dto"
)

type SignInUseCase struct {
	Repository ports.IRepository
	Hasher     ports.IHasher
	JWT        ports.IJWT
}

var AccountNotFound error = errors.New("account not found")
var InCorrectCredentials error = errors.New("email or password are incorrect")

func NewSignInUseCase(repository ports.IRepository, hasher ports.IHasher, jwt ports.IJWT) *SignInUseCase {
	return &SignInUseCase{
		Repository: repository,
		Hasher:     hasher,
		JWT:        jwt,
	}
}

func (u *SignInUseCase) SignIn(userDto dto.SignInDTO) (ports.Token, error) {
	user, err := u.Repository.FindByEmail(userDto.Email)

	if user.Email == "" {
		return "", AccountNotFound
	}

	if err != nil {
		return "", err
	}

	isPasswordMatch, err := u.Hasher.Compare(user.Password, userDto.Password)

	if err != nil {
		return "", InCorrectCredentials
	}

	if !isPasswordMatch {
		return "", InCorrectCredentials
	}

	user.Password = ""

	return u.JWT.Generate(user), nil
}
