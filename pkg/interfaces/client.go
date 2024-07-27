package interfaces

import (
	"context"

	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
)

type ClientRepository interface {
	CreateClient(ctx context.Context, client *entities.Client) (*entities.Client, error)
	GetClientByCPF(ctx context.Context, cpf string) (*entities.Client, error)
}
