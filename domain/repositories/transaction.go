package repositories

type SaveTransactionRepository interface {
	Save(id string, account string, amount float64, status string, errorMessage string) error
}
