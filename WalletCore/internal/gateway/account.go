package gateway

import "github.com.br/luciano340/FC_microservices/WalletCore/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}
