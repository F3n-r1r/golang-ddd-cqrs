package valueobjects

import (
	"errors"
	"fmt"
)

type FullName struct {
	firstName string
	lastName  string
}

func NewFullName(firstName string, lastName string) (FullName, error) {
	if firstName == "" {
		return FullName{}, errors.New("...")
	}

	if lastName == "" {
		return FullName{}, errors.New("...")
	}

	f := FullName{
		firstName: firstName,
		lastName:  lastName,
	}

	return f, nil
}

func (f FullName) String() string {
	return fmt.Sprintf("%s %s", f.firstName, f.lastName)
}
