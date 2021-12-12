package main

import (
	"database/sql"
	"encoding/json"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jhonpedro/imersaofc5/gateway/adapter/broker/kafka"
	"github.com/jhonpedro/imersaofc5/gateway/adapter/factory"
	"github.com/jhonpedro/imersaofc5/gateway/adapter/presenter/transaction"
	uniqueId "github.com/jhonpedro/imersaofc5/gateway/adapter/services"
	"github.com/jhonpedro/imersaofc5/gateway/usecase/process_transaction"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "main.db")

	if err != nil {
		log.Fatal(err)
	}

	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)

	repository := repositoryFactory.CreateTransactionRepository()

	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
	}

	kafkaPresenter := transaction.NewKafkaPresenter()

	producer := kafka.NewKafkaProducer(configMapProducer, kafkaPresenter)

	var msgChannel = make(chan *ckafka.Message)

	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
		"client.id":         "goapp",
		"group.id":          "goapp",
	}

	topics := []string{
		"transactions",
	}

	consumer := kafka.NewConsumer(configMapConsumer, topics)

	go consumer.Consume(msgChannel)

	usecase := process_transaction.NewProcessTransaction(repository, uniqueId.NewUniqueIdentifierService(), producer, "transactions_result")

	for msg := range msgChannel {
		var input process_transaction.TransactionInputDto

		json.Unmarshal(msg.Value, &input)

		usecase.Execute(input)
	}

}
