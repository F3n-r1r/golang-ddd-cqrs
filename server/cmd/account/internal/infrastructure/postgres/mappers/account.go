package mappers

import (
	"errors"
	"golang-ddd-cqrs/cmd/account/internal/domain/model/aggregates"
	"golang-ddd-cqrs/cmd/account/internal/domain/model/valueobjects"
	"golang-ddd-cqrs/cmd/account/internal/infrastructure/models"
)

// Map from domain Account to DB Account
func ToDBAccount(a *aggregates.Account) *models.Account {
	account := &models.Account{
		ID:       a.GetID(),
		Email:    a.Email.String(),
		Password: a.Password.String(),
	}

	return account
}

func ToDomainAccount(a *models.Account) (*aggregates.Account, error) {
	account, err := aggregates.NewAccount()
	if err != nil {
		return &aggregates.Account{}, errors.New("...")
	}

	account.Hydrate("test")

	id, err := valueobjects.ParseID(a.ID)
	if err != nil {
		return &aggregates.Account{}, errors.New("...")
	}

	fullName, err := valueobjects.NewFullName(a.FirstName, a.LastName)
	if err != nil {
		return &aggregates.Account{}, errors.New("...")
	}

	email, err := valueobjects.NewEmail(a.Email)
	if err != nil {
		return &aggregates.Account{}, errors.New("...")
	}

	address, err := valueobjects.NewAddress(a.City, a.Zipcode, a.Street, a.Number)
	if err != nil {
		return &aggregates.Account{}, errors.New("...")
	}

	// Hydration refers to the process of filling an object with data. An object which has not yet been hydrated has been instantiated and represents an entity that does have data, but the data has not yet been loaded into the object. This is something that is done for performance reasons.
	return &aggregates.Account.Hydrate(id, fullName, email, address), nil
}

func ToDomainAccounts(u []*Account) ([]*aggregates.Account, error) {
	dcs := make([]*aggregates.Account, len(u))
	for k, v := range u {
		dc, err := ToDomainAccount(v)
		if err != nil {
			return nil, err
		}
		dcs[k] = dc
	}
	return dcs, nil
}
