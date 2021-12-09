package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jhonpedro/imersaofc5/gateway/adapter/presenter"
)

type Producer struct {
	ConfigMap *ckafka.ConfigMap
	Presenter presenter.Presenter
}

func NewKafkaProducer(configMap *ckafka.ConfigMap, presenter presenter.Presenter) *Producer {
	return &Producer{ConfigMap: configMap, Presenter: presenter}
}

func (p *Producer) Publish(msg interface{}, key []byte, topic string) error {
	producer, errNewProducer := ckafka.NewProducer(p.ConfigMap)

	if errNewProducer != nil {
		return errNewProducer
	}

	errBind := p.Presenter.Bind(msg)

	if errBind != nil {
		return errBind
	}

	presenterMsg, errShow := p.Presenter.Show()

	if errShow != nil {
		return errShow
	}

	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{
			Topic:     &topic,
			Partition: ckafka.PartitionAny,
		},
		Value: presenterMsg,
		Key:   key,
	}

	errProducedMessage := producer.Produce(message, nil)

	if errProducedMessage != nil {
		return errProducedMessage
	}

	return nil
}
