package main

import (
	"database/sql"
	"encoding/json"
	"github.com/GustavoEklund/payment-gateway-api/application/controller"
	"github.com/GustavoEklund/payment-gateway-api/domain/use_cases/process_transaction"
	"github.com/GustavoEklund/payment-gateway-api/infra/brokers/kafka"
	"github.com/GustavoEklund/payment-gateway-api/infra/repositories/sqlite"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	transactionRepositoryFactory := sqlite.NewRepositoryDatabaseFactory(db)
	transactionRepository := transactionRepositoryFactory.Make()
	producerConfigMap := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}
	transactionController := controller.NewTransactionController()
	kafkaProducer := kafka.NewKafkaProducer(producerConfigMap, transactionController)
	msgChan := make(chan *ckafka.Message)
	producerConfigMap = &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "go-app",
		"group.id":          "go-app",
	}
	topics := []string{"transactions"}
	consumer := kafka.NewConsumer(producerConfigMap, topics)
	go consumer.Consume(msgChan)
	useCase := process_transaction.NewProcessTransaction(transactionRepository, kafkaProducer, "transactions_result")
	for msg := range msgChan {
		var input process_transaction.TransactionInput
		json.Unmarshal(msg.Value, &input)
		useCase.Perform(input)
	}
}
