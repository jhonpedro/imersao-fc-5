package factory

import "github.com/jhonpedro/imersaofc5/gateway/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
