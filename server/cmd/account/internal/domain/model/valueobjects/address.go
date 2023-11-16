package valueobjects

import (
	"errors"
	"fmt"
)

type Address struct {
	city    string
	zipcode int
	street  string
	number  int
}

func NewAddress(city string, zipcode int, street string, number int) (Address, error) {
	if city == "" {
		return Address{}, errors.New("...")
	}

	if zipcode == 0 {
		return Address{}, errors.New("...")
	}

	if street == "" {
		return Address{}, errors.New("...")
	}

	if number == 0 {
		return Address{}, errors.New("...")
	}

	a := Address{
		city:    city,
		zipcode: zipcode,
		street:  street,
		number:  number,
	}

	return a, nil
}

func (a Address) String() string {
	return fmt.Sprintf("%s %s, %s %s", a.street, a.number, a.zipcode, a.city)
}
