package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/mfritschdotgo/techchallengefase2/pkg/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBPaymentRepository struct {
	Collection *mongo.Collection
}

func NewMongoDBPaymentRepository(db *mongo.Database) interfaces.PaymentRepository {
	return &MongoDBPaymentRepository{
		Collection: db.Collection("orders"),
	}
}

func (r *MongoDBPaymentRepository) UpdatePayment(ctx context.Context, id uuid.UUID, status int, description string) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": status, "status_description": description}}
	_, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
