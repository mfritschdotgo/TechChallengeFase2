package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/mfritschdotgo/techchallengefase2/internal/adapters/presenters"
	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
	"github.com/mfritschdotgo/techchallengefase2/internal/domain/usecases"
	"github.com/mfritschdotgo/techchallengefase2/pkg/dto"
)

type CategoryController struct {
	usecases  *usecases.Category
	presenter presenters.CategoryPresenter
}

func NewCategoryController(s *usecases.Category, p presenters.CategoryPresenter) *CategoryController {
	return &CategoryController{
		usecases:  s,
		presenter: p,
	}
}

// CreateCategory adds a new category to the store
// @Summary Add a new category
// @Description Adds a new category to the database with the given details.
// @Tags categories
// @Accept json
// @Produce json
// @Param		request	body		dto.CreateCategoryRequest	true	"Category creation details"
// @Success 201 {object} entities.Category "Successfully created Category"
// @Failure 400 "Bad request if the Category data is invalid"
// @Failure 500 "Internal server error if there is a problem on the server side"
// @Router /categories [post]
func (c *CategoryController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var categoryDto dto.CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&categoryDto); err != nil {
		c.presenter.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	category, err := c.usecases.CreateCategory(ctx, categoryDto)
	if err != nil {
		c.presenter.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.presenter.JSON(w, category, http.StatusCreated)
}

// ReplaceCategory replace an existing category with the provided details.
// @Summary Replace an existing category
// @Description Replace category details in the database by ID.
// @Tags categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Param		request	body		dto.CreateCategoryRequest	true	"Category object that needs to be updated"
// @Success 200 {object} entities.Category "Category successfully updated"
// @Failure 400 {string} string "Invalid input, Object is invalid"
// @Failure 404 {string} string "Category not found"
// @Failure 500 {string} string "Internal server error"
// @Router /categories/{id} [put]
func (c *CategoryController) ReplaceCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	if id == "" {
		c.presenter.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	var category entities.Category

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		c.presenter.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := c.usecases.ReplaceCategory(ctx, id, &category)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.presenter.Error(w, "Category not found", http.StatusNotFound)
		} else {
			c.presenter.Error(w, "Error replacing category", http.StatusInternalServerError)
		}
		return
	}

	c.presenter.JSON(w, response, http.StatusOK)
}

// UpdateCategory update an existing category with the provided details.
// @Summary Update an existing category
// @Description Update category details in the database by ID.
// @Tags categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Param		request	body		dto.CreateCategoryRequest	true	"Category object that needs to be updated"
// @Success 200 {object} entities.Category "Category successfully updated"
// @Failure 400 {string} string "Invalid input, Object is invalid"
// @Failure 404 {string} string "Category not found"
// @Failure 500 {string} string "Internal server error"
// @Router /categories/{id} [patch]
func (c *CategoryController) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	if id == "" {
		c.presenter.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	var category entities.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		c.presenter.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := c.usecases.UpdateCategory(ctx, id, &category)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.presenter.Error(w, "Category not found", http.StatusNotFound)
		} else {
			c.presenter.Error(w, "Error updating category", http.StatusInternalServerError)
		}
		return
	}

	c.presenter.JSON(w, response, http.StatusOK)
}

// GetCategoryByID retrieves a category by its ID
// @Summary Get a category
// @Description Retrieves details of a category based on its unique ID.
// @Tags categories
// @Accept json
// @Produce json
// @Param id path string true "category ID"
// @Success 200 {object} entities.Category "Successfully retrieved the category details"
// @Failure 400 "Bad request if the ID is not provided or invalid"
// @Failure 404 "Product not found if the ID does not match any category"
// @Router /categories/{id} [get]
func (c *CategoryController) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	if id == "" {
		c.presenter.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	category, err := c.usecases.GetCategoryByID(ctx, id)
	if err != nil {
		c.presenter.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	c.presenter.JSON(w, category, http.StatusOK)
}

// GetCategories retrieves a list of categories
// @Summary List categories
// @Description Retrieves a paginated list of categories
// @Tags categories
// @Accept json
// @Produce json
// @Param page query int false "Page number for pagination" default(1)
// @Param pageSize query int false "Number of categories per page" default(10)
// @Success 200 {array} entities.Category "Successfully retrieved list of categories"
// @Failure 500 "Internal server error if there is a problem on the server side"
// @Router /categories [get]
func (c *CategoryController) GetCategories(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil || size <= 0 {
		size = 10
	}

	categories, err := c.usecases.GetCategories(ctx, page, size)
	if err != nil {
		c.presenter.Error(w, "Failed to retrieve categories", http.StatusInternalServerError)
		return
	}

	c.presenter.JSON(w, categories, http.StatusOK)
}

// DeleteCategory deletes a category by its ID
// @Summary Delete a category
// @Description Deletes a category based on its unique ID and returns a success message.
// @Tags categories
// @Accept json
// @Produce json
// @Param id path string true "category ID"
// @Success 200 {object} map[string]string "Message indicating successful deletion"
// @Failure 400 "Bad request if the ID is not provided or is invalid"
// @Failure 404 "category not found if the ID does not match any category"
// @Failure 500 "Internal server error if there is a problem deleting the category"
// @Router /categories/{id} [delete]
func (c *CategoryController) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	if id == "" {
		c.presenter.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	if err := c.usecases.DeleteCategory(ctx, id); err != nil {
		c.presenter.Error(w, "Category not found or error deleting category", http.StatusNotFound)
		return
	}

	response := map[string]string{"message": "Category with ID " + id + " deleted successfully."}
	c.presenter.JSON(w, response, http.StatusOK)
}
