package process_transaction

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mock_broker "github.com/jhonpedro/imersaofc5/gateway/adapter/broker/mock"
	"github.com/jhonpedro/imersaofc5/gateway/domain/entities"
	mock_repository "github.com/jhonpedro/imersaofc5/gateway/domain/repository/mock"
	mock_services "github.com/jhonpedro/imersaofc5/gateway/services/mock"
	"github.com/stretchr/testify/assert"
)

func TestProcessTransaction_ExecuteInvalidCreditCard(t *testing.T) {
	input := TransactionInputDto{

		Id:                        "uuid",
		AccountId:                 "1",
		CreditCardNumber:          "34111111111111",
		CreditCardName:            "João Pedro",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    200,
	}

	expectedOutput := TransactionOutputDto{
		EvaluationId: "uuid",
		Id:           "uuid",
		Status:       entities.REJECTED,
		ErrorMessage: "invalid credit card number",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().Insert("uuid", input.Id, input.AccountId, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).Return(nil)

	uniqueIdentifierMock := mock_services.NewMockUniqueIdentifierService(ctrl)
	uniqueIdentifierMock.EXPECT().Generate().Return("uuid")

	producerMock := mock_broker.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.Id), "transaction_result")

	usecase := NewProcessTransaction(repositoryMock, uniqueIdentifierMock, producerMock, "transaction_result")

	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestProcessTransaction_ExecuteRejectTransaction(t *testing.T) {
	input := TransactionInputDto{
		Id:                        "uuid",
		AccountId:                 "1",
		CreditCardNumber:          "341111111111111",
		CreditCardName:            "João Pedro",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    2000,
	}

	expectedOutput := TransactionOutputDto{
		EvaluationId: "uuid",
		Id:           "uuid",
		Status:       entities.REJECTED,
		ErrorMessage: "you do not have limit for this transaction",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().Insert("uuid", input.Id, input.AccountId, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).Return(nil)

	uniqueIdentifierMock := mock_services.NewMockUniqueIdentifierService(ctrl)
	uniqueIdentifierMock.EXPECT().Generate().Return("uuid")

	producerMock := mock_broker.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.Id), "transaction_result")

	usecase := NewProcessTransaction(repositoryMock, uniqueIdentifierMock, producerMock, "transaction_result")

	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestProcessTransaction_ExecuteApprovedTransaction(t *testing.T) {
	input := TransactionInputDto{
		Id:                        "uuid",
		AccountId:                 "1",
		CreditCardNumber:          "341111111111111",
		CreditCardName:            "João Pedro",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    900,
	}

	expectedOutput := TransactionOutputDto{
		EvaluationId: "uuid",
		Id:           "uuid",
		Status:       entities.APPROVED,
		ErrorMessage: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().Insert("uuid", input.Id, input.AccountId, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).Return(nil)

	uniqueIdentifierMock := mock_services.NewMockUniqueIdentifierService(ctrl)
	uniqueIdentifierMock.EXPECT().Generate().Return("uuid")

	producerMock := mock_broker.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.Id), "transaction_result").Return(nil)

	usecase := NewProcessTransaction(repositoryMock, uniqueIdentifierMock, producerMock, "transaction_result")

	output, err := usecase.Execute(input)

	fmt.Println(err)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
