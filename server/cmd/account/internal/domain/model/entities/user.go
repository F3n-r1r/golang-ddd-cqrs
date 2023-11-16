package entities

import (
	"errors"
	"golang-ddd-cqrs/cmd/account/internal/domain/model/valueobjects"
)

type User struct {
	ID       valueobjects.ID
	Email    valueobjects.Email
	Password valueobjects.Password
	FullName valueobjects.FullName
	Address  valueobjects.Address
}

// User constructor - instantiate new base user object with required fields
func NewUser(e string, p string, f string, l string) (*User, error) {
	id := valueobjects.NewID()

	email, err := valueobjects.NewEmail(e)
	if err != nil {
		return &User{}, errors.New("...")
	}

	password, err := valueobjects.NewPassword(p)
	if err != nil {
		return &User{}, errors.New("...")
	}

	fullName, err := valueobjects.NewFullName(f, l)
	if err != nil {
		return &User{}, errors.New("...")
	}

	return &User{
		ID:       id,
		Email:    email,
		Password: password,
		FullName: fullName,
	}, nil
}

func HydrateUser(id valueobjects.ID, email valueobjects.Email, password valueobjects.Password, fullname valueobjects.FullName, address valueobjects.Address) *User {
	return &User{
		ID:       id,
		Email:    email,
		Password: password,
		FullName: fullname,
	}
}

// func (u *User) ParseUserID(s string) (uuid.UUID, error) {
// 	id, err := uuid.Parse(s)
// 	if err != nil {
// 		return uuid.UUID{}, err
// 	}

// 	return id, nil
// }
