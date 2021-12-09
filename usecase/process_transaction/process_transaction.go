package process_transaction

import (
	"github.com/jhonpedro/imersaofc5/gateway/adapter/broker"
	"github.com/jhonpedro/imersaofc5/gateway/domain/entities"
	"github.com/jhonpedro/imersaofc5/gateway/domain/repository"
	"github.com/jhonpedro/imersaofc5/gateway/services"
)

type ProcessTransaction struct {
	repository       repository.TransactionRepository
	uniqueIdentifier services.UniqueIdentifierService
	producer         broker.ProducerInterface
	topic            string
}

func NewProcessTransaction(repository repository.TransactionRepository, uniqueIdentifier services.UniqueIdentifierService, producer broker.ProducerInterface, topic string) *ProcessTransaction {
	return &ProcessTransaction{repository, uniqueIdentifier, producer, topic}
}

func (p *ProcessTransaction) Execute(input TransactionInputDto) (TransactionOutputDto, error) {

	id := p.uniqueIdentifier.Generate()

	transaction := entities.NewTransaction(id, input.AccountId, input.Amount)
	cc, invalidCcError := entities.NewCreditCard(input.CreditCardNumber, input.CreditCardName, input.CreditCardExpirationMonth, input.CreditCardExpirationYear, input.CreditCardCVV)

	if invalidCcError != nil {
		return p.rejectTransaction(input, id, invalidCcError)
	}

	transaction.SetCreditCard(*cc)
	isTransactionValid := transaction.IsValid()

	if isTransactionValid != nil {
		return p.rejectTransaction(input, id, isTransactionValid)
	}

	return p.aproveTransaction(input, id)
}

func (p *ProcessTransaction) rejectTransaction(input TransactionInputDto, id string, err error) (TransactionOutputDto, error) {
	repositoryInsertError := p.repository.Insert(id, input.AccountId, input.Amount, entities.REJECTED, err.Error())

	if repositoryInsertError != nil {
		return TransactionOutputDto{}, repositoryInsertError
	}

	output := TransactionOutputDto{
		Id:           id,
		Status:       entities.REJECTED,
		ErrorMessage: err.Error(),
	}

	errPub := p.publish(output, []byte(input.Id))

	if errPub != nil {
		return TransactionOutputDto{}, errPub
	}

	return output, nil
}

func (p *ProcessTransaction) aproveTransaction(input TransactionInputDto, id string) (TransactionOutputDto, error) {
	repositoryInsertError := p.repository.Insert(id, input.AccountId, input.Amount, entities.APPROVED, "")

	if repositoryInsertError != nil {
		return TransactionOutputDto{}, repositoryInsertError
	}

	output := TransactionOutputDto{
		Id:           id,
		Status:       entities.APPROVED,
		ErrorMessage: "",
	}

	errPub := p.publish(output, []byte(input.Id))

	if errPub != nil {
		return TransactionOutputDto{}, errPub
	}

	return output, nil
}

func (p *ProcessTransaction) publish(output TransactionOutputDto, key []byte) error {
	err := p.producer.Publish(output, key, p.topic)

	if err != nil {
		return err
	}

	return nil
}
