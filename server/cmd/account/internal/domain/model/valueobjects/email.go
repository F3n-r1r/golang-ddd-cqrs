package valueobjects

import (
	"errors"
	"regexp"
)

type Email struct {
	Email string
}

var (
	ErrInvalidEmail = errors.New("email is invalid")
	ErrEmptyEmail   = errors.New("email cannot be empty")
)

func NewEmail(email string) (Email, error) {
	if email == "" {
		return Email{}, ErrEmptyEmail
	}

	e := Email{
		Email: email,
	}

	if err := e.Validate(); err != nil {
		return Email{}, ErrEmptyEmail
	}

	return e, nil
}

func (e *Email) Validate() error {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !emailRegex.MatchString(e.Email) {
		return ErrInvalidEmail
	}

	return nil
}

func (e *Email) String() string {
	return e.Email
}
