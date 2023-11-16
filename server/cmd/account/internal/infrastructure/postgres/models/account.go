package models

import (
	"github.com/google/uuid"
)

type Account struct {
	ID       uuid.UUID
	Email    string
	Password string
	City     string
	Zipcode  int
	Street   string
	Number   int
	// Created_at time.Time
	// Updated_at time.Time
}
