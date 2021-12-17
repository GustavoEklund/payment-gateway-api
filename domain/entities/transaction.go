package entities

import "errors"

const (
	STATUS_REJECTED = "rejected"
	STATUS_APPROVED = "approved"
)

type Transaction struct {
	ID           string
	AccountID    string
	Amount       float64
	CreditCard   CreditCard
	Status       string
	ErrorMessage string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) IsValid() error {
	if t.Amount > 1000 {
		return errors.New("your limit is not enough for this transaction")
	}
	if t.Amount < 1 {
		return errors.New("the amount must be greater than 1")
	}
	return nil
}

func (t *Transaction) SetCreditCard(creditCard CreditCard) {
	t.CreditCard = creditCard
}
