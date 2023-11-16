package services

import (
	"golang-ddd-cqrs/cmd/account/internal/domain/model/aggregates"
	"golang-ddd-cqrs/cmd/account/internal/domain/model/valueobjects"
)

// Interface what the infrastructure repository needs to implement.
type UserRepository interface {
	FindByEmail(email string) (*aggregates.Account, error)
}

type AccountService struct {
	repo UserRepository
}

type createAccountInput struct {
	email                string
	password             string
	passwordConfirmation string
}

func NewAccountDomainService(r UserRepository) *AccountService {
	return &AccountService{
		repo: r,
	}
}

func (s *AccountService) CreateAccount(input createAccountInput) (*aggregates.Account, error) {
	// Verify that the email does not exists
	if _, err := s.repo.FindByEmail(input.email); err != nil {
		return &aggregates.Account{}, err
	}

	// Note: Move password and email creation/validation to NewAccount function?

	// Create, validate & hash password value object
	password, err := valueobjects.NewPassword(input.password)
	if err != nil {
		return &aggregates.Account{}, err
	}

	if err := password.ConfirmMatch(input.passwordConfirmation); err != nil {
		return &aggregates.Account{}, err
	}

	if err := password.Validate(); err != nil {
		return &aggregates.Account{}, err
	}

	if err := password.Hash(); err != nil {
		return &aggregates.Account{}, err
	}

	// Create & validate email value object
	email, err := valueobjects.NewEmail(input.email)
	if err != nil {
		return &aggregates.Account{}, err
	}

	if err := email.Validate(); err != nil {
		return &aggregates.Account{}, err
	}

	// Create account aggregate & attach value objects
	account, err := aggregates.NewAccount()
	if err != nil {
		return &aggregates.Account{}, err
	}

	account.SetPassword(password.String())
	account.SetEmail(email.String())

	return account, nil
}
