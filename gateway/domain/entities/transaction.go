package entities

import "errors"

const (
	REJECTED = "rejected"
	APPROVED = "approved"
)

type Transaction struct {
	EvaluationId string
	Id           string
	AccountId    string
	Amount       float64
	Status       string
	ErrorMessage string
	CreditCard   CreditCard
}

func NewTransaction(evaluation_id string, id string, accountId string, amount float64) *Transaction {
	return &Transaction{
		evaluation_id,
		id,
		accountId,
		amount,
		"",
		"",
		CreditCard{},
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

func (t *Transaction) SetCreditCard(card CreditCard) {
	t.CreditCard = card
}
