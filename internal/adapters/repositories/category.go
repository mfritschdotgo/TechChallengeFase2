package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
	"github.com/mfritschdotgo/techchallengefase2/pkg/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBCategoryRepository struct {
	collection *mongo.Collection
}

func NewMongoDBCategoryRepository(db *mongo.Database) interfaces.CategoryRepository {
	return &MongoDBCategoryRepository{
		collection: db.Collection("categories"),
	}
}

func (r *MongoDBCategoryRepository) CreateCategory(ctx context.Context, category *entities.Category) (*entities.Category, error) {
	_, err := r.collection.InsertOne(ctx, category)
	if err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}
	return category, nil
}

func (r *MongoDBCategoryRepository) GetCategoryByID(ctx context.Context, id string) (*entities.Category, error) {
	var category entities.Category

	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}

	binaryUUID := primitive.Binary{
		Subtype: 0x00,
		Data:    uuid[:],
	}

	filter := bson.M{"_id": binaryUUID}
	err = r.collection.FindOne(ctx, filter).Decode(&category)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("category not found")
		}
		return nil, fmt.Errorf("failed to get category by ID: %w", err)
	}
	return &category, nil
}

func (r *MongoDBCategoryRepository) ReplaceCategory(ctx context.Context, category *entities.Category) (*entities.Category, error) {
	uuid, err := uuid.Parse(category.ID.String())
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}

	binaryUUID := primitive.Binary{
		Subtype: 0x00,
		Data:    uuid[:],
	}

	filter := bson.M{"_id": binaryUUID}
	update := bson.M{"$set": category}

	result, err := r.collection.ReplaceOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to replace category: %w", err)
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("category not found")
	}

	return category, nil
}

func (r *MongoDBCategoryRepository) UpdateCategory(ctx context.Context, category *entities.Category) (*entities.Category, error) {
	uuid, err := uuid.Parse(category.ID.String())
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
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

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update category: %w", err)
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("category not found")
	}

	return category, nil
}

func (r *MongoDBCategoryRepository) DeleteCategory(ctx context.Context, id string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid ID format: %w", err)
	}

	binaryUUID := primitive.Binary{
		Subtype: 0x00,
		Data:    uuid[:],
	}

	filter := bson.M{"_id": binaryUUID}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("category not found")
	}

	return nil
}

func (r *MongoDBCategoryRepository) GetCategories(ctx context.Context, page, limit int) ([]entities.Category, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	var categories []entities.Category
	opts := options.Find().SetSkip(int64((page - 1) * limit)).SetLimit(int64(limit))
	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var category entities.Category
		if err := cursor.Decode(&category); err != nil {
			return nil, fmt.Errorf("failed to decode category: %w", err)
		}
		categories = append(categories, category)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return categories, nil
}
