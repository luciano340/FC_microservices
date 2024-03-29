package gateway

import "github.com.br/luciano340/FC_microservices/WalletCore/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
