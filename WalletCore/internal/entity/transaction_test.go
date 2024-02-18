package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransacation(t *testing.T) {
	client1, _ := NewClient("Luciano", "contato@lucianoromao.com.br")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Taynara", "contato@taynara.com.br")
	account2 := NewAccount(client2)

	account1.Credit(2000)
	account2.Credit(2000)

	transaction, err := NewTransaction(account1, account2, 200)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 2200.0, account2.Balance)
	assert.Equal(t, 1800.0, account1.Balance)

}

func TestCreateTransactionWithInsuficientBalance(t *testing.T) {
	client1, _ := NewClient("Luciano", "contato@lucianoromao.com.br")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Taynara", "contato@taynara.com.br")
	account2 := NewAccount(client2)

	account1.Credit(2000)
	account2.Credit(2000)

	transaction, err := NewTransaction(account1, account2, 3000)
	assert.NotNil(t, err)
	assert.Error(t, err, "Insufficient founds")
	assert.Nil(t, transaction)
	assert.Equal(t, 2000.0, account1.Balance)
	assert.Equal(t, 2000.0, account2.Balance)

	transaction2, err2 := NewTransaction(account1, account2, 0)
	assert.NotNil(t, err2)
	assert.Nil(t, transaction2)
	assert.Error(t, err2, "Amount must be grater than zero")

}
