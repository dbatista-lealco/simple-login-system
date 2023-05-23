package ports

import "github.com/dylanbatar/simple-login-system/internal/domain"

type IRepository interface {
	Save(user domain.User) error
}
