package process_transaction

type CreditCard struct {
	Number          string `json:"number"`
	Name            string `json:"name"`
	ExpirationMonth int    `json:"expiration_month"`
	ExpirationYear  int    `json:"expiration_year"`
	Cvv             int    `json:"cvv"`
}

type TransactionInput struct {
	ID         string     `json:"id"`
	AccountID  string     `json:"account_id"`
	CreditCard CreditCard `json:"credit_card"`
	Amount     float64    `json:"amount"`
}

type TransactionOutput struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}
