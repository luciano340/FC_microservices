package gateway

import "github.com.br/luciano340/FC_microservices/WalletCore/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
