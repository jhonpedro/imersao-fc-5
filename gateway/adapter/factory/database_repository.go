package factory

import (
	"database/sql"

	repo "github.com/jhonpedro/imersaofc5/gateway/adapter/repository"
	"github.com/jhonpedro/imersaofc5/gateway/domain/repository"
)

type RepositoryDatabaseFactory struct {
	DB *sql.DB
}

func NewRepositoryDatabaseFactory(db *sql.DB) *RepositoryDatabaseFactory {
	return &RepositoryDatabaseFactory{DB: db}
}

func (r RepositoryDatabaseFactory) CreateTransactionRepository() repository.TransactionRepository {
	return repo.NewTransactionRepositoryDb(r.DB)
}
