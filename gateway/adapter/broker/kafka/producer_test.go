package kafka

import (
	"testing"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jhonpedro/imersaofc5/gateway/adapter/presenter/transaction"
	"github.com/jhonpedro/imersaofc5/gateway/domain/entities"
	"github.com/jhonpedro/imersaofc5/gateway/usecase/process_transaction"
	"github.com/stretchr/testify/assert"
)

func TestProducerPublish(t *testing.T) {
	expectOutput := process_transaction.TransactionOutputDto{
		Id:           "uuid",
		Status:       entities.REJECTED,
		ErrorMessage: "you dont have limit for this transaction",
	}

	// outputJson, _ := json.Marshal(expectOutput)

	configMap := ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
	}
	producer := NewKafkaProducer(&configMap, transaction.NewKafkaPresenter())

	err := producer.Publish(expectOutput, []byte("1"), "test")

	assert.Nil(t, err)
}
