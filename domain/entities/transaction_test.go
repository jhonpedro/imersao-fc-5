package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionWithAmountGreaterThan1000(t *testing.T) {
	transaction := NewTransaction("1", "1", 2000)

	isTransactionValid := transaction.IsValid()

	assert.Error(t, isTransactionValid)
	assert.Equal(t, "you do not have limit for this transaction", isTransactionValid.Error())
}

func TestTransactionWithAmoutLessThan1(t *testing.T) {
	transaction := NewTransaction("1", "1", 0)

	isTransactionValid := transaction.IsValid()

	assert.Error(t, isTransactionValid)
	assert.Equal(t, "the amount must be greater than 1", isTransactionValid.Error())
}
