package create_transaction

import (
	"github.com.br/luciano340/FC_microservices/WalletCore/internal/entity"
	"github.com.br/luciano340/FC_microservices/WalletCore/internal/gateway"
)

type CreateTransactionInputDTO struct {
	AccountIdFrom string
	AccountIDTo   string
	Amount        float64
}

type CreateTransactionOutputDTO struct {
	ID string
}

type CreateTransactionUseCase struct {
	TransacationGateway gateway.TransactionGateway
	AccountGateway      gateway.AccountGateway
}

func NewCreateTransactionUseCase(transacationGateway gateway.TransactionGateway, accountGateway gateway.AccountGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransacationGateway: transacationGateway,
		AccountGateway:      accountGateway,
	}
}

func (c *CreateTransactionUseCase) Execute(input CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	accountFrom, err := c.AccountGateway.FindByID(input.AccountIdFrom)
	if err != nil {
		return nil, err
	}

	accountTo, err := c.AccountGateway.FindByID(input.AccountIDTo)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}
	err = c.TransacationGateway.Create(transaction)
	if err != nil {
		return nil, err
	}
	return &CreateTransactionOutputDTO{ID: transaction.ID}, nil
}
