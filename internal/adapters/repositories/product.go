package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
	"github.com/mfritschdotgo/techchallengefase2/pkg/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBProductRepository struct {
	Collection *mongo.Collection
}

func NewMongoDBProductRepository(db *mongo.Database) interfaces.ProductRepository {
	return &MongoDBProductRepository{Collection: db.Collection("products")}
}

func (r *MongoDBProductRepository) CreateProduct(ctx context.Context, product *entities.Product) (*entities.Product, error) {
	_, err := r.Collection.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *MongoDBProductRepository) GetProductByID(ctx context.Context, id string) (*entities.Product, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	binaryUUID := primitive.Binary{
		Subtype: 0x00,
		Data:    uuid[:],
	}

	var product entities.Product
	err = r.Collection.FindOne(ctx, bson.M{"_id": binaryUUID}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *MongoDBProductRepository) ReplaceProduct(ctx context.Context, product *entities.Product) (*entities.Product, error) {
	uuid, err := uuid.Parse(product.ID.String())
	if err != nil {
		return nil, err
	}

	binaryUUID := primitive.Binary{
		Subtype: 0x00,
		Data:    uuid[:],
	}

	filter := bson.M{"_id": binaryUUID}
	update := bson.M{"$set": product}
	_, err = r.Collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *MongoDBProductRepository) UpdateProduct(ctx context.Context, product *entities.Product) (*entities.Product, error) {
	uuid, err := uuid.Parse(product.ID.String())
	if err != nil {
		return nil, err
	}

	binaryUUID := primitive.Binary{
		Subtype: 0x00,
		Data:    uuid[:],
	}

	filter := bson.M{"_id": binaryUUID}
	update := bson.M{"$set": product}

	if product.CategoryId.String() != "" {
		update["$set"].(bson.M)["category_id"] = product.CategoryId
	}

	if product.Name != "" {
		update["$set"].(bson.M)["name"] = product.Name
	}

	if product.Description != "" {
		update["$set"].(bson.M)["description"] = product.Description
	}

	if product.Price != 0 {
		update["$set"].(bson.M)["price"] = product.Price
	}

	if !product.UpdatedAt.IsZero() {
		update["$set"].(bson.M)["updated_at"] = product.UpdatedAt
	}

	_, err = r.Collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *MongoDBProductRepository) DeleteProduct(ctx context.Context, id string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	binaryUUID := primitive.Binary{
		Subtype: 0x00,
		Data:    uuid[:],
	}

	_, err = r.Collection.DeleteOne(ctx, bson.M{"_id": binaryUUID})

	if err != nil {
		return err
	}

	return nil
}

func (r *MongoDBProductRepository) GetProducts(ctx context.Context, categoryId string, page, limit int) ([]entities.Product, error) {
	var products []entities.Product
	opts := options.Find().SetSkip(int64((page - 1) * limit)).SetLimit(int64(limit))
	filter := bson.M{}

	if categoryId != "" {

		uuid, err := uuid.Parse(categoryId)
		if err != nil {
			return nil, err
		}

		binaryUUID := primitive.Binary{
			Subtype: 0x00,
			Data:    uuid[:],
		}

		filter["category_id"] = binaryUUID
	}

	cursor, err := r.Collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product entities.Product
		if err = cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
