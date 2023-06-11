package domain

type PasswordReset struct {
	Code   string
	Email  string
	Active bool
}

func NewPasswordReset(code, email string) PasswordReset {
	// TODO: currently a code is active if is the last otp code generate by the user. we need that code has a TTL
	return PasswordReset{
		Code:   code,
		Email:  email,
		Active: true,
	}
}
