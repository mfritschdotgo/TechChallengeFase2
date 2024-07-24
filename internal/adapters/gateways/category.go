package gateways

import (
	"context"

	"github.com/google/uuid"
	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CategoryRepository struct {
	Collection *mongo.Collection
}

func NewCategoryRepository(db *mongo.Database) *CategoryRepository {
	return &CategoryRepository{
		Collection: db.Collection("categories"),
	}
}

func (cr *CategoryRepository) CreateCategory(ctx context.Context, category *entities.Category) (*entities.Category, error) {
	_, err := cr.Collection.InsertOne(ctx, category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (cr *CategoryRepository) GetCategoryByID(ctx context.Context, id string) (*entities.Category, error) {
	var category entities.Category

	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	binaryUUID := primitive.Binary{
		Subtype: 0x00,
		Data:    uuid[:],
	}

	filter := bson.M{"_id": binaryUUID}
	err = cr.Collection.FindOne(ctx, filter).Decode(&category)

	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (cr *CategoryRepository) ReplaceCategory(ctx context.Context, category *entities.Category) (*entities.Category, error) {
	uuid, err := uuid.Parse(category.ID.String())
	if err != nil {
		return nil, err
	}

	binaryUUID := primitive.Binary{
		Subtype: 0x00,
		Data:    uuid[:],
	}

	filter := bson.M{"_id": binaryUUID}
	update := bson.M{"$set": category}

	_, err = cr.Collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}
	return category, nil
}

func (cr *CategoryRepository) UpdateCategory(ctx context.Context, category *entities.Category) (*entities.Category, error) {
	uuid, err := uuid.Parse(category.ID.String())
	if err != nil {
		return nil, err
	}

	binaryUUID := primitive.Binary{
		Subtype: 0x00,
		Data:    uuid[:],
	}

	filter := bson.M{"_id": binaryUUID}
	update := bson.M{"$set": bson.M{}}

	if category.Name != "" {
		update["$set"].(bson.M)["name"] = category.Name
	}
	if category.Description != "" {
		update["$set"].(bson.M)["description"] = category.Description
	}

	_, err = cr.Collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}
	return category, nil
}

func (cr *CategoryRepository) DeleteCategory(ctx context.Context, id string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	binaryUUID := primitive.Binary{
		Subtype: 0x00,
		Data:    uuid[:],
	}

	filter := bson.M{"_id": binaryUUID}
	_, err = cr.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (cr *CategoryRepository) GetCategories(ctx context.Context, page, limit int) ([]entities.Category, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	var categories []entities.Category
	opts := options.Find().SetSkip(int64((page - 1) * limit)).SetLimit(int64(limit))
	cursor, err := cr.Collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var category entities.Category
		if err = cursor.Decode(&category); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
