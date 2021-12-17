package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreditCard_Number(t *testing.T) {
	_, err := NewCreditCard("6666666666666666", "Any Full Name", 12, 2024, 123)
	assert.Error(t, err)
	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard("5357204502621242", "Any Full Name", 12, 2024, 123)
	assert.Nil(t, err)
}

func TestCreditCard_ExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("5357204502621242", "Any Full Name", 13, 2024, 123)
	assert.Error(t, err)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("5357204502621242", "Any Full Name", 0, 2024, 123)
	assert.Error(t, err)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("5357204502621242", "Any Full Name", 11, 2024, 123)
	assert.Nil(t, err)
}

func TestCreditCard_ExpirationYear(t *testing.T) {
	lastYear := time.Now().AddDate(-1, 0, 0)
	_, err := NewCreditCard("5357204502621242", "Any Full Name", 12, lastYear.Year(), 123)
	assert.Equal(t, "invalid expiration year", err.Error())

	_, err = NewCreditCard("5357204502621242", "Any Full Name", 12, lastYear.Year()+12, 123)
	assert.Equal(t, "invalid expiration year", err.Error())

	_, err = NewCreditCard("5357204502621242", "Any Full Name", 12, 2024, 123)
	assert.Nil(t, err)
}

func TestCreditCard_Cvv(t *testing.T) {
	_, err := NewCreditCard("5357204502621242", "Any Full Name", 12, 2024, 1000)
	assert.Equal(t, "invalid cvv", err.Error())

	_, err = NewCreditCard("5357204502621242", "Any Full Name", 12, 2024, 99)
	assert.Equal(t, "invalid cvv", err.Error())

	_, err = NewCreditCard("5357204502621242", "Any Full Name", 12, 2024, 321)
	assert.Nil(t, err)
}
