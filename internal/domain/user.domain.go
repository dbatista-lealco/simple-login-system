package domain

import (
	"github.com/dylanbatar/simple-login-system/internal/domain/utils"
	"time"
)

type User struct {
	Name             string
	Email            string
	Password         string
	Image            string
	Role             string
	Active           bool
	LastActivity     []LastActivity
	VerifyCationCode string
	IsVerify         bool
}

type LastActivity struct {
	Address string
	Ip      string
	Date    time.Time
}

func NewUser(name, email, password string) (*User, error) {
	return &User{
		Name:             name,
		Email:            email,
		Password:         password,
		Image:            "",
		Role:             "user",
		Active:           true,
		LastActivity:     []LastActivity{},
		VerifyCationCode: utils.GenerateVerifyCode(5),
		IsVerify:         false,
	}, nil
}
