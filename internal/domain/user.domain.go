package domain

import (
	"github.com/dylanbatar/simple-login-system/internal/domain/utils"
	"time"
)

type User struct {
	Name             string         `bson:"name"`
	Email            string         `bson:"email"`
	Password         string         `bson:"password"`
	Image            string         `bson:"image"`
	Role             string         `bson:"role"`
	Active           bool           `bson:"active"`
	LastActivity     []LastActivity `bson:"lastActivity"`
	VerifyCationCode string         `bson:"verifyCationCode"`
	IsVerify         bool           `bson:"isVerify"`
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
