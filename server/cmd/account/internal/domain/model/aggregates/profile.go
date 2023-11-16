package aggregates

import "github.com/google/uuid"

type Auth struct {
	ID     uuid.UUID
	UserID uuid.UUID
}
