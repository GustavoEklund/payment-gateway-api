package process_transaction

import (
	"github.com/GustavoEklund/payment-gateway-api/domain/entities"
	"github.com/GustavoEklund/payment-gateway-api/domain/repositories"
	"github.com/GustavoEklund/payment-gateway-api/infra/brokers"
)

type ProcessTransaction struct {
	SaveTransactionRepository repositories.SaveTransactionRepository
	Producer                  brokers.Producer
	Topic                     string
}

func NewProcessTransaction(
	SaveTransactionRepository repositories.SaveTransactionRepository,
	producer brokers.Producer,
	topic string) *ProcessTransaction {
	return &ProcessTransaction{
		SaveTransactionRepository: SaveTransactionRepository,
		Producer:                  producer,
		Topic:                     topic,
	}
}

func (p *ProcessTransaction) Perform(input TransactionInput) (TransactionOutput, error) {
	transaction := entities.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount
	creditCard, creditCardErr := entities.NewCreditCard(input.CreditCard.Number, input.CreditCard.Name, input.CreditCard.ExpirationMonth, input.CreditCard.ExpirationYear, input.CreditCard.Cvv)
	if creditCardErr != nil {
		return p.rejectTransaction(transaction, creditCardErr)
	}
	transaction.SetCreditCard(*creditCard)
	transactionErr := transaction.IsValid()
	if transactionErr != nil {
		return p.rejectTransaction(transaction, transactionErr)
	}
	return p.approveTransaction(transaction)
}

func (p *ProcessTransaction) approveTransaction(transaction *entities.Transaction) (TransactionOutput, error) {
	err := p.SaveTransactionRepository.Save(transaction.ID, transaction.AccountID, transaction.Amount, entities.STATUS_APPROVED, "")
	if err != nil {
		return TransactionOutput{}, err
	}
	output := TransactionOutput{
		ID:           transaction.ID,
		Status:       entities.STATUS_APPROVED,
		ErrorMessage: "",
	}
	err = p.Publish(output, []byte(transaction.ID))
	if err != nil {
		return TransactionOutput{}, err
	}
	return output, nil
}

func (p *ProcessTransaction) rejectTransaction(transaction *entities.Transaction, transactionErr error) (TransactionOutput, error) {
	err := p.SaveTransactionRepository.Save(transaction.ID, transaction.AccountID, transaction.Amount, entities.STATUS_REJECTED, transactionErr.Error())
	if err != nil {
		return TransactionOutput{}, err
	}
	output := TransactionOutput{
		ID:           transaction.ID,
		Status:       entities.STATUS_REJECTED,
		ErrorMessage: transactionErr.Error(),
	}
	err = p.Publish(output, []byte(transaction.ID))
	if err != nil {
		return TransactionOutput{}, err
	}
	return output, nil
}

func (p *ProcessTransaction) Publish(output TransactionOutput, key []byte) error {
	err := p.Producer.Publish(output, key, p.Topic)
	if err != nil {
		return err
	}
	return nil
}
