package query

import (
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID
	Email string
}

type Listing struct {
	ID    uuid.UUID
	Title string
}

type GetUserListingsDto struct {
	User     User
	Listings []Listing
}

type GetUserDto struct {
	User User
}

type GetUsersDto struct {
	Users []User
}
