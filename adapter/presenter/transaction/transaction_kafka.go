package transaction

import (
	"encoding/json"

	"github.com/jhonpedro/imersaofc5/gateway/usecase/process_transaction"
)

type KafkaPresenter struct {
	Id           string `json:"id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func NewKafkaPresenter() *KafkaPresenter {
	return &KafkaPresenter{}
}

func (k *KafkaPresenter) Bind(input interface{}) error {
	k.Id = input.(process_transaction.TransactionOutputDto).Id
	k.Status = input.(process_transaction.TransactionOutputDto).Status
	k.ErrorMessage = input.(process_transaction.TransactionOutputDto).ErrorMessage

	return nil
}

func (k *KafkaPresenter) Show() ([]byte, error) {
	j, err := json.Marshal(k)

	if err != nil {
		return nil, err
	}

	return j, err
}
