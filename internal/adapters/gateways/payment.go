package gateways

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentRepository struct {
	Collection *mongo.Collection
}

func NewPaymentRepository(db *mongo.Database) *PaymentRepository {
	return &PaymentRepository{
		Collection: db.Collection("orders"),
	}
}

func (pr *PaymentRepository) UpdatePayment(ctx context.Context, id uuid.UUID, status int, description string) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": status, "status_description": description}}
	_, err := pr.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
