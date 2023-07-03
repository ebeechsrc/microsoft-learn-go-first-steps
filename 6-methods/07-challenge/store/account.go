package store

type Account struct {
	name Name
}

func (a *Account) Name() Name {
	return a.name
}

func (a *Account) ChangeName(newName Name) {
	a.name = newName
}
