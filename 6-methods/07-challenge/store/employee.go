package store

import "errors"

type Employee struct {
	Account
	credits float64
}

var ErrInvalidCreditAmount = errors.New("Invalid credit amount")
var ErrInsufficientCredits = errors.New("Insufficient credits")

func (e *Employee) Credits() float64 {
	return e.credits
}

func (e *Employee) AddCredits(amount float64) (float64, error) {
	if amount < 0 {
		return 0, ErrInvalidCreditAmount
	}
	e.credits += amount
	return e.credits, nil
}

func (e *Employee) RemoveCredits(amount float64) (float64, error) {
	if amount < 0 {
		return 0, ErrInvalidCreditAmount
	}
	if amount > e.credits {
		return 0, ErrInsufficientCredits
	}
	e.credits -= amount
	return e.credits, nil
}
