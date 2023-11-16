package valueobjects

import (
	"errors"

	"github.com/google/uuid"
)

type ID struct {
	id uuid.UUID
}

func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)

	if err != nil {
		return ID{}, errors.New("...")
	}

	c := ID{
		id: id,
	}

	return c, nil
}

func NewID() ID {
	return ID{
		id: uuid.New(),
	}
}

func MustParseID(
	id string,
) ID {
	c, err := ParseID(id)
	if err != nil {
		panic(err)
	}
	return c
}

func (c ID) Id() uuid.UUID {
	return c.id
}

func (c ID) IsZero() bool {
	return c == ID{}
}

func (c ID) String() string {
	return c.id.String()
}
