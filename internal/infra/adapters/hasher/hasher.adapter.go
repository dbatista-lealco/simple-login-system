package hasher

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type HasherAdapter struct{}

func NewHasherAdapter() HasherAdapter {
	return HasherAdapter{}
}

func (adapter HasherAdapter) Hash(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("error hashing password %s", err.Error())
	}

	return string(hashPassword), nil
}

func (adapter HasherAdapter) Compare(hashPassoword, password string) (bool, error) {
	panic("")
}
