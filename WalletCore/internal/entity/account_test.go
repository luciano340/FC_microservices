package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("Luciano Romão", "contato@lucianoromao.com.br")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateAccountWithNiLClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("Luciano Romão", "contato@lucianoromao.com.br")
	account := NewAccount(client)
	account.Credit(100)
	assert.Equal(t, float64(100), account.Balance)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("Luciano Romão", "contato@lucianoromao.com.br")
	account := NewAccount(client)
	account.Credit(200)
	account.Debit(120)
	assert.Equal(t, float64(80), account.Balance)
}
