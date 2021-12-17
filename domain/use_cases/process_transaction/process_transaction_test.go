package process_transaction

import (
	"github.com/GustavoEklund/payment-gateway-api/domain/entities"
	mocks "github.com/GustavoEklund/payment-gateway-api/domain/repositories/mocks"
	mock_brokers "github.com/GustavoEklund/payment-gateway-api/infra/brokers/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestProcessTransaction_PerformInvalidCreditCard(t *testing.T) {
	input := TransactionInput{
		ID:        "1",
		AccountID: "1",
		CreditCard: CreditCard{
			Number:          "6666666666666666",
			Name:            "Any Full Name",
			ExpirationMonth: 12,
			ExpirationYear:  time.Now().Year() + 1,
			Cvv:             123,
		},
		Amount: 200,
	}
	expectedOutput := TransactionOutput{
		ID:           "1",
		Status:       entities.STATUS_REJECTED,
		ErrorMessage: "invalid credit card number",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	saveTransactionRepositorySpy := mocks.NewMockSaveTransactionRepository(ctrl)
	saveTransactionRepositorySpy.EXPECT().
		Save(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)
	producerSpy := mock_brokers.NewMockProducer(ctrl)
	producerSpy.EXPECT().
		Publish(expectedOutput, []byte(input.ID), "transactions_result").
		Return(nil)

	sut := NewProcessTransaction(saveTransactionRepositorySpy, producerSpy, "transactions_result")
	output, err := sut.Perform(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestProcessTransaction_PerformRejectedTransaction(t *testing.T) {
	input := TransactionInput{
		ID:        "1",
		AccountID: "1",
		CreditCard: CreditCard{
			Number:          "5357204502621242",
			Name:            "Any Full Name",
			ExpirationMonth: 12,
			ExpirationYear:  time.Now().Year() + 1,
			Cvv:             123,
		},
		Amount: 1200,
	}
	expectedOutput := TransactionOutput{
		ID:           "1",
		Status:       entities.STATUS_REJECTED,
		ErrorMessage: "your limit is not enough for this transaction",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	saveTransactionRepositorySpy := mocks.NewMockSaveTransactionRepository(ctrl)
	saveTransactionRepositorySpy.EXPECT().
		Save(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)
	producerSpy := mock_brokers.NewMockProducer(ctrl)
	producerSpy.EXPECT().
		Publish(expectedOutput, []byte(input.ID), "transactions_result").
		Return(nil)

	sut := NewProcessTransaction(saveTransactionRepositorySpy, producerSpy, "transactions_result")
	output, err := sut.Perform(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestProcessTransaction_PerformApprovedTransaction(t *testing.T) {
	input := TransactionInput{
		ID:        "1",
		AccountID: "1",
		CreditCard: CreditCard{
			Number:          "5357204502621242",
			Name:            "Any Full Name",
			ExpirationMonth: 12,
			ExpirationYear:  time.Now().Year() + 1,
			Cvv:             123,
		},
		Amount: 200,
	}
	expectedOutput := TransactionOutput{
		ID:           "1",
		Status:       entities.STATUS_APPROVED,
		ErrorMessage: "",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	saveTransactionRepositorySpy := mocks.NewMockSaveTransactionRepository(ctrl)
	saveTransactionRepositorySpy.EXPECT().
		Save(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)
	producerSpy := mock_brokers.NewMockProducer(ctrl)
	producerSpy.EXPECT().
		Publish(expectedOutput, []byte(input.ID), "transactions_result").
		Return(nil)

	sut := NewProcessTransaction(saveTransactionRepositorySpy, producerSpy, "transactions_result")
	output, err := sut.Perform(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
