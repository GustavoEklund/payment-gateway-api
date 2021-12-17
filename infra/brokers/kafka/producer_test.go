package kafka

import (
	"github.com/GustavoEklund/payment-gateway-api/application/controller"
	"github.com/GustavoEklund/payment-gateway-api/domain/entities"
	"github.com/GustavoEklund/payment-gateway-api/domain/use_cases/process_transaction"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKafkaProducerPublish(t *testing.T) {
	expectedOutput := process_transaction.TransactionOutput{
		ID:           "1",
		Status:       entities.STATUS_REJECTED,
		ErrorMessage: "your limit is not enough for this transaction",
	}
	//jsonOutput, _ := json.Marshal(expectedOutput)

	configMap := ckafka.ConfigMap{"test.mock.num.brokers": 3}
	producer := NewKafkaProducer(&configMap, controller.NewTransactionController())
	err := producer.Publish(expectedOutput, []byte("1"), "test")

	assert.Nil(t, err)
}
