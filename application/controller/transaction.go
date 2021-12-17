package controller

import (
	"encoding/json"
	"github.com/GustavoEklund/payment-gateway-api/domain/use_cases/process_transaction"
)

type TransactionController struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func NewTransactionController() *TransactionController {
	return &TransactionController{}
}

func (t *TransactionController) Bind(input interface{}) error {
	t.ID = input.(process_transaction.TransactionOutput).ID
	t.Status = input.(process_transaction.TransactionOutput).Status
	t.ErrorMessage = input.(process_transaction.TransactionOutput).ErrorMessage
	return nil
}

func (t *TransactionController) Handle() ([]byte, error) {
	requestBody, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return requestBody, nil
}
