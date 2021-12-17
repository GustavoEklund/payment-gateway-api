package kafka

import (
	"github.com/GustavoEklund/payment-gateway-api/application/controller"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	ConfigMap  *ckafka.ConfigMap
	Controller controller.Controller
}

func NewKafkaProducer(configMap *ckafka.ConfigMap, controller controller.Controller) *Producer {
	return &Producer{ConfigMap: configMap, Controller: controller}
}

func (p *Producer) Publish(msg interface{}, key []byte, topic string) error {
	producer, err := ckafka.NewProducer(p.ConfigMap)
	if err != nil {
		return err
	}
	err = p.Controller.Bind(msg)
	if err != nil {
		return err
	}
	presenterMessage, err := p.Controller.Handle()
	if err != nil {
		return err
	}
	message := ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          presenterMessage,
		Key:            key,
	}
	err = producer.Produce(&message, nil)
	if err != nil {
		return err
	}
	return nil
}
