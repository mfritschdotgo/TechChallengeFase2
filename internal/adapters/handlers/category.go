package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/mfritschdotgo/techchallengefase2/internal/adapters/controllers"
	"github.com/mfritschdotgo/techchallengefase2/internal/adapters/presenters"
	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
	"github.com/mfritschdotgo/techchallengefase2/pkg/dto"
)

type CategoryHandler struct {
	controller *controllers.CategoryController
	presenter  presenters.CategoryPresenter
}

func NewCategoryHandler(c *controllers.CategoryController, p presenters.CategoryPresenter) *CategoryHandler {
	return &CategoryHandler{
		controller: c,
		presenter:  p,
	}
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var categoryDto dto.CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&categoryDto); err != nil {
		h.presenter.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	category, err := h.controller.CreateCategory(r.Context(), categoryDto)
	if err != nil {
		h.presenter.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.presenter.JSON(w, category, http.StatusCreated)
}

func (h *CategoryHandler) ReplaceCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		h.presenter.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	var category entities.Category

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		h.presenter.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.controller.ReplaceCategory(r.Context(), id, &category)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.presenter.Error(w, "Category not found", http.StatusNotFound)
		} else {
			h.presenter.Error(w, "Error replacing category", http.StatusInternalServerError)
		}
		return
	}

	h.presenter.JSON(w, response, http.StatusOK)
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		h.presenter.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	var category entities.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		h.presenter.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.controller.UpdateCategory(r.Context(), id, &category)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.presenter.Error(w, "Category not found", http.StatusNotFound)
		} else {
			h.presenter.Error(w, "Error updating category", http.StatusInternalServerError)
		}
		return
	}

	h.presenter.JSON(w, response, http.StatusOK)
}

func (h *CategoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		h.presenter.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	category, err := h.controller.GetCategoryByID(r.Context(), id)
	if err != nil {
		h.presenter.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	h.presenter.JSON(w, category, http.StatusOK)
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil || size <= 0 {
		size = 10
	}

	categories, err := h.controller.GetCategories(r.Context(), page, size)
	if err != nil {
		h.presenter.Error(w, "Failed to retrieve categories", http.StatusInternalServerError)
		return
	}

	h.presenter.JSON(w, categories, http.StatusOK)
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		h.presenter.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	if err := h.controller.DeleteCategory(r.Context(), id); err != nil {
		h.presenter.Error(w, "Category not found or error deleting category", http.StatusNotFound)
		return
	}

	response := map[string]string{"message": "Category with ID " + id + " deleted successfully."}
	h.presenter.JSON(w, response, http.StatusOK)
}
