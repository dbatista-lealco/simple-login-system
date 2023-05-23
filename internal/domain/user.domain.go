package domain

import "time"

type User struct {
	Name         string
	Email        string
	Password     string
	Image        string
	Role         string
	Active       bool
	LastActivity []LastActivity
}

type LastActivity struct {
	Address string
	Ip      string
	Date    time.Time
}

func NewUser(name, email, password string) (*User, error) {
	return &User{
		Name:         name,
		Email:        email,
		Password:     password,
		Image:        "",
		Role:         "user",
		Active:       true,
		LastActivity: []LastActivity{},
	}, nil
}
