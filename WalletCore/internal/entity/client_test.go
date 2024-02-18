package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("Luciano", "contato@lucianoromao.com.br")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Luciano", client.Name)
	assert.Equal(t, "contato@lucianoromao.com.br", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("Luciano Romão", "contato@lucianoromao.com.br")
	err := client.Update("Luciano Souza", "contato@gmail.com")
	assert.Nil(t, err)
	assert.Equal(t, "Luciano Souza", client.Name)
	assert.Equal(t, "contato@gmail.com", client.Email)
}

func TestUpdateClientWhenArgsAreInvalid(t *testing.T) {
	client, _ := NewClient("Luciano Romão", "contato@lucianoromao.com.br")
	err := client.Update("", "contato@gmail.com")
	assert.Error(t, err, "Name is required")
}

func TestAccountToClient(t *testing.T) {
	client, _ := NewClient("Luciano Romão", "contato@lucianoromao.com.br")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))

	client2, _ := NewClient("Taynara Gomes", "contato@gmail.com")
	account2 := NewAccount(client2)
	err2 := client.AddAccount(account2)
	assert.Error(t, err2, "Account does not belong to client")
}
