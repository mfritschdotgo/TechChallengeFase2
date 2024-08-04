package gateways

import (
	"context"

	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
	"github.com/mfritschdotgo/techchallengefase2/pkg/interfaces"
)

type ClientGateway struct {
	repo interfaces.ClientRepository
}

func NewClientGateway(repo interfaces.ClientRepository) *ClientGateway {
	return &ClientGateway{
		repo: repo,
	}
}

func (g *ClientGateway) CreateClient(ctx context.Context, client *entities.Client) (*entities.Client, error) {
	return g.repo.CreateClient(ctx, client)
}

func (g *ClientGateway) GetClientByCPF(ctx context.Context, cpf string) (*entities.Client, error) {
	return g.repo.GetClientByCPF(ctx, cpf)
}
