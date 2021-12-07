package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumberIsInvalid(t *testing.T) {
	_, err := NewCreditCard("40000000000", "Joãozinho Rei das Gambiarras", 12, 2024, 123)

	assert.Equal(t, "invalid credit card number", err.Error())
}

func TestCreditCardMonthIsInvalid(t *testing.T) {
	creditCardNumber := "341111111111111"

	_, err := NewCreditCard(creditCardNumber, "Joãozinho Rei das Gambiarras", 13, 2024, 123)

	assert.Equal(t, "invalid expiration month", err.Error())
}

func TestCreditCardMonthIsValid(t *testing.T) {
	creditCardNumber := "341111111111111"

	_, err := NewCreditCard(creditCardNumber, "Joãozinho Rei das Gambiarras", 12, 2024, 123)

	assert.Nil(t, err)
}

func TestCreditCardNumberIsValid(t *testing.T) {
	creditCardNumber := "341111111111111"

	_, err := NewCreditCard(creditCardNumber, "Joãozinho Rei das Gambiarras", 12, 2024, 123)

	assert.Nil(t, err)
}

func TestCreditCardYearIsInvalid(t *testing.T) {
	creditCardNumber := "341111111111111"

	lastYear := time.Now().AddDate(-1, 0, 0)

	_, error := NewCreditCard(creditCardNumber, "Joãozinho Rei das Gambiarras", 12, lastYear.Year(), 123)

	assert.Equal(t, "invalid expiration year", error.Error())
}
