package query

import (
	"golang-ddd-cqrs/cmd/user/internal/domain"
)

// check for map methods -> https://github.com/quintans/go-clean-ddd/blob/master/internal/app/query/customer.go

func mapToSingle(user *domain.User) User {
	return User{
		ID:    user.GetID(),
		Email: user.Email.String(),
	}
}

func mapToList(users []*domain.User) []User {
	result := make([]User, len(users))

	for i, o := range users {
		result[i] = mapToSingle(o)
	}

	return result
}
