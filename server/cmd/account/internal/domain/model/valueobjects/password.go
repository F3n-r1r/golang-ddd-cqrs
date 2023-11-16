package valueobjects

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	Password string
}

var (
	ErrPasswordConfirmationNoMatch = errors.New("passwords do not match")
	ErrPasswordLength              = errors.New("password must be at least 8 characters long")
	ErrPasswordHash                = errors.New("password hashing failed")
)

func NewPassword(p string) (Password, error) {
	if p == "" {
		return Password{}, ErrEmptyEmail
	}

	password := Password{
		Password: p,
	}

	if err := password.Validate(); err != nil {
		return Password{}, ErrPasswordLength
	}

	if err := password.Hash(); err != nil {
		return Password{}, err
	}

	return password, nil
}

func (e *Password) Validate() error {
	// Check length
	if len(e.Password) < 8 {
		return ErrPasswordLength
	}

	// Check else?

	return nil
}

func (e *Password) ConfirmMatch(passwordConfirmation string) error {
	if e.Password != passwordConfirmation {
		return ErrPasswordConfirmationNoMatch
	}
	return nil
}

func (e *Password) Hash() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost)

	if err != nil {
		return ErrPasswordHash
	}

	e.Password = string(hashedPassword)

	return nil
}

func (e *Password) String() string {
	return e.Password
}
