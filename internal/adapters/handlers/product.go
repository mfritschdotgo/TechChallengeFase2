package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/mfritschdotgo/techchallengefase2/internal/adapters/controllers"
	"github.com/mfritschdotgo/techchallengefase2/pkg/dto"
	"github.com/mfritschdotgo/techchallengefase2/pkg/interfaces"
)

type ProductHandler struct {
	controller *controllers.ProductController
	presenter  interfaces.ProductPresenter
}

func NewProductHandler(c *controllers.ProductController, p interfaces.ProductPresenter) *ProductHandler {
	return &ProductHandler{
		controller: c,
		presenter:  p,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var productDto dto.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&productDto); err != nil {
		h.presenter.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	product, err := h.controller.CreateProduct(ctx, productDto)
	if err != nil {
		h.presenter.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.presenter.JSON(w, product, http.StatusCreated)
}

func (h *ProductHandler) ReplaceProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	if id == "" {
		h.presenter.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var productDto dto.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&productDto); err != nil {
		h.presenter.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	product, err := h.controller.ReplaceProduct(ctx, id, productDto)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.presenter.Error(w, err.Error(), http.StatusNotFound)
		} else {
			h.presenter.Error(w, "Error updating product", http.StatusInternalServerError)
		}
		return
	}

	h.presenter.JSON(w, product, http.StatusOK)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	if id == "" {
		h.presenter.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var productDto dto.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&productDto); err != nil {
		h.presenter.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	product, err := h.controller.UpdateProduct(ctx, id, productDto)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.presenter.Error(w, err.Error(), http.StatusNotFound)
		} else {
			h.presenter.Error(w, "Error updating product", http.StatusInternalServerError)
		}
		return
	}

	h.presenter.JSON(w, product, http.StatusOK)
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	if id == "" {
		h.presenter.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.controller.GetProductByID(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.presenter.Error(w, "Product not found", http.StatusNotFound)
		} else {
			h.presenter.Error(w, "Error retrieving product", http.StatusInternalServerError)
		}
		return
	}

	h.presenter.JSON(w, product, http.StatusOK)
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil || size <= 0 {
		size = 10
	}

	category := r.URL.Query().Get("category")

	products, err := h.controller.GetProducts(ctx, category, page, size)
	if err != nil {
		h.presenter.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.presenter.JSON(w, products, http.StatusOK)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	if id == "" {
		h.presenter.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err := h.controller.DeleteProduct(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.presenter.Error(w, "Product not found", http.StatusNotFound)
		} else {
			h.presenter.Error(w, "Error deleting product", http.StatusInternalServerError)
		}
		return
	}

	response := map[string]string{"message": "Product with ID " + id + " deleted successfully."}
	h.presenter.JSON(w, response, http.StatusOK)
}
