package dto

type CreateCategoryRequest struct {
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
}
