package aggregates

import (
	"golang-ddd-cqrs/cmd/account/internal/domain/model/entities"
)

// Account is a aggregate that combines all entities/value objects needed to represent a Account
type Account struct {
	// user is the root entity of an account which means the User.ID is the main identifier for this aggregate
	User *entities.User
}

// Account constructor - instantiate new base account object with required fields
func NewAccount(email string, password string, firstname string, lastname string) (*Account, error) {
	user, err := entities.NewUser(email, password, firstname, lastname)
	if err != nil {
		return &Account{}, err
	}

	return &Account{
		User: user,
	}, nil
}

func HydrateAccount() *Account {
	user := entities.HydrateUser()

	return &Account{
		User: user,
	}
}

// func (a *Account) Hydrate(t string) *Account {
// 	user := entities.NewUser()
// 	user.Hydrate()
// 	account := &Account{}
// 	return account
// }

// func (a *Account) GetID() uuid.UUID {
// 	return a.User.ID
// }

// func (a *Account) SetPassword(password string) {
// 	a.Password.Password = string(password)
// }

// func (a *Account) SetEmail(email string) {
// 	a.Email.Email = string(email)
// }
