package store

import "fmt"

type Name struct {
	Forename, Surname string
}

func (name Name) String() string {
	return fmt.Sprintf("%v %v", name.Forename, name.Surname)
}
