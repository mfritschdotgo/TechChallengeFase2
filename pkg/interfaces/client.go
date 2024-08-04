package interfaces

import (
	"context"

	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
	"github.com/mfritschdotgo/techchallengefase2/pkg/dto"
)

type ClientRepository interface {
	CreateClient(ctx context.Context, client *entities.Client) (*entities.Client, error)
	GetClientByCPF(ctx context.Context, cpf string) (*entities.Client, error)
}

type ClientUseCase interface {
	CreateClient(ctx context.Context, dto dto.CreateClientRequest) (*entities.Client, error)
	GetClientByCPF(ctx context.Context, cpf string) (*entities.Client, error)
}

type ClientGateway interface {
	CreateClient(ctx context.Context, client *entities.Client) (*entities.Client, error)
	GetClientByCPF(ctx context.Context, cpf string) (*entities.Client, error)
}
