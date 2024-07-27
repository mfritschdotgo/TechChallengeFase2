package gateways

import (
	"context"

	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClientRepository struct {
	Collection *mongo.Collection
}

func NewClientRepository(db *mongo.Database) *ClientRepository {
	return &ClientRepository{
		Collection: db.Collection("clients"),
	}
}

func (r *ClientRepository) CreateClient(ctx context.Context, client *entities.Client) (*entities.Client, error) {
	_, err := r.Collection.InsertOne(ctx, client)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (r *ClientRepository) GetClientByCPF(ctx context.Context, cpf string) (*entities.Client, error) {
	var client entities.Client
	err := r.Collection.FindOne(ctx, bson.M{"cpf": cpf}).Decode(&client)
	if err != nil {
		return nil, err
	}
	return &client, nil
}
