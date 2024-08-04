package usecases

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
	"github.com/mfritschdotgo/techchallengefase2/pkg/dto"
	"github.com/mfritschdotgo/techchallengefase2/pkg/interfaces"
)

type Client struct {
	gateway interfaces.ClientGateway
}

func NewClient(gateway interfaces.ClientGateway) interfaces.ClientUseCase {
	return &Client{
		gateway: gateway,
	}
}

func (s *Client) CreateClient(ctx context.Context, dto dto.CreateClientRequest) (*entities.Client, error) {
	if !isValidCPF(dto.Cpf) {
		return nil, fmt.Errorf("invalid CPF format")
	}

	if !isValidEmail(dto.Mail) {
		return nil, fmt.Errorf("invalid email format")
	}

	client, err := entities.NewClient(dto.Name, dto.Cpf, dto.Mail)
	if err != nil {
		return nil, err
	}

	_, err = s.gateway.CreateClient(ctx, client)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (s *Client) GetClientByCPF(ctx context.Context, cpf string) (*entities.Client, error) {
	client, err := s.gateway.GetClientByCPF(ctx, cpf)
	if err != nil {
		return nil, fmt.Errorf("client not found: %w", err)
	}

	return client, nil
}

func isValidEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}

func isValidCPF(cpf string) bool {
	cpf = regexp.MustCompile(`[^0-9]+`).ReplaceAllString(cpf, "")
	if len(cpf) != 11 {
		return false
	}

	var sum int
	var remainder int

	for i := 1; i <= 9; i++ {
		digit, _ := strconv.Atoi(string(cpf[i-1]))
		sum += digit * (11 - i)
	}
	remainder = (sum * 10) % 11

	if remainder == 10 || remainder == 11 {
		remainder = 0
	}
	if remainder != int(cpf[9]-'0') {
		return false
	}

	sum = 0
	for i := 1; i <= 10; i++ {
		digit, _ := strconv.Atoi(string(cpf[i-1]))
		sum += digit * (12 - i)
	}
	remainder = (sum * 10) % 11

	if remainder == 10 || remainder == 11 {
		remainder = 0
	}
	if remainder != int(cpf[10]-'0') {
		return false
	}

	return true
}
