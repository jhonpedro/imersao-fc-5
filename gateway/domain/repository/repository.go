package repository

type TransactionRepository interface {
	Insert(evaluation_id string, id string, accountId string, amount float64, status string, errorMessage string) error
}
