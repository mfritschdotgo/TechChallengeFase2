package repositories

import (
	"context"

	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
	"github.com/mfritschdotgo/techchallengefase2/pkg/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBClientRepository struct {
	collection *mongo.Collection
}

func NewMongoDBClientRepository(db *mongo.Database) interfaces.ClientRepository {
	return &MongoDBClientRepository{
		collection: db.Collection("clients"),
	}
}

func (r *MongoDBClientRepository) CreateClient(ctx context.Context, client *entities.Client) (*entities.Client, error) {
	_, err := r.collection.InsertOne(ctx, client)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (r *MongoDBClientRepository) GetClientByCPF(ctx context.Context, cpf string) (*entities.Client, error) {
	var client entities.Client
	err := r.collection.FindOne(ctx, bson.M{"cpf": cpf}).Decode(&client)
	if err != nil {
		return nil, err
	}
	return &client, nil
}
