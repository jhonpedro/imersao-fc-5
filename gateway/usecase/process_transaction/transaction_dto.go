package process_transaction

type TransactionInputDto struct {
	Id                        string  `json:"id"`
	AccountId                 string  `json:"account_id"`
	CreditCardNumber          string  `json:"credit_card_number"`
	CreditCardName            string  `json:"credit_card_name"`
	CreditCardExpirationMonth int     `json:"credit_card_expiration_month"`
	CreditCardExpirationYear  int     `json:"credit_card_expiration_year"`
	CreditCardCVV             int     `json:"credit_card_cvv"`
	Amount                    float64 `json:"amount"`
}

type TransactionOutputDto struct {
	EvaluationId string `json:"evaluation_id"`
	Id           string `json:"id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}
