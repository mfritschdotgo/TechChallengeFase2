package gateways

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderRepository struct {
	Collection        *mongo.Collection
	CounterCollection *mongo.Collection
}

func NewOrderRepository(db *mongo.Database) *OrderRepository {
	return &OrderRepository{
		Collection:        db.Collection("orders"),
		CounterCollection: db.Collection("counters"),
	}
}

func (pr *OrderRepository) CreateOrder(ctx context.Context, order *entities.Order) (*entities.Order, error) {
	counter, err := pr.getCurrentCounter(ctx)
	if err != nil {
		return nil, err
	}
	order.Order = counter

	_, err = pr.Collection.InsertOne(ctx, order)

	if err != nil {
		return nil, err
	}
	return order, nil
}

func (pr *OrderRepository) GetOrders(ctx context.Context, page, limit int) ([]entities.Order, error) {
	var orders []entities.Order
	opts := options.Find().
		SetSkip(int64((page - 1) * limit)).
		SetLimit(int64(limit)).
		SetSort(bson.D{{Key: "status", Value: -1}, {Key: "CreatedAt", Value: 1}})

	filter := bson.M{"status": bson.M{"$ne": 4}}

	cursor, err := pr.Collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var order entities.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (pr *OrderRepository) GetOrderByID(ctx context.Context, id string) (*entities.Order, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	binaryUUID := primitive.Binary{
		Subtype: 0x00,
		Data:    uuid[:],
	}

	var order entities.Order
	err = pr.Collection.FindOne(ctx, bson.M{"_id": binaryUUID}).Decode(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (pr *OrderRepository) SetStatus(ctx context.Context, id uuid.UUID, status int, description string) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": status, "status_description": description}}
	_, err := pr.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (pr *OrderRepository) getCurrentCounter(ctx context.Context) (int, error) {
	today := time.Now().Format("2006-01-02")
	var result struct {
		Counter int
	}
	err := pr.CounterCollection.FindOneAndUpdate(
		ctx,
		bson.M{"date": today},
		bson.M{"$inc": bson.M{"counter": 1}},
		options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After),
	).Decode(&result)

	if err != nil {
		return 0, err
	}

	return result.Counter, nil
}
