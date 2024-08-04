package controllers

import (
	"context"

	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
	"github.com/mfritschdotgo/techchallengefase2/pkg/dto"
	"github.com/mfritschdotgo/techchallengefase2/pkg/interfaces"
)

type ClientController struct {
	usecase interfaces.ClientUseCase
}

func NewClientController(u interfaces.ClientUseCase) *ClientController {
	return &ClientController{
		usecase: u,
	}
}

func (c *ClientController) CreateClient(ctx context.Context, clientDto dto.CreateClientRequest) (*entities.Client, error) {
	return c.usecase.CreateClient(ctx, clientDto)
}

func (c *ClientController) GetClientByCPF(ctx context.Context, cpf string) (*entities.Client, error) {
	return c.usecase.GetClientByCPF(ctx, cpf)
}
