package repository

import (
	"os"
	"testing"

	"github.com/jhonpedro/fullcycle-go/adapter/repository/fixture"
	"github.com/jhonpedro/imersaofc5/gateway/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestTransactionRepositoryDb_Insert(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)

	repository := NewTransactionRepositoryDb(db)

	err := repository.Insert("1", "1", 2, entities.APPROVED, "")

	assert.Nil(t, err)

}
