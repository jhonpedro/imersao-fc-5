package entities

import "errors"

type Transaction struct {
	Id           string
	AccountId    string
	Amount       float64
	Status       string
	ErrorMessage string
}

func NewTransaction(id string, accountId string, amount float64) *Transaction {
	return &Transaction{
		id,
		accountId,
		amount,
		"",
		"",
	}
}

func (t *Transaction) IsValid() error {
	if t.Amount > 1000 {
		return errors.New("you do not have limit for this transaction")
	}
	if t.Amount < 1 {
		return errors.New("the amount must be greater than 1")
	}

	return nil
}
