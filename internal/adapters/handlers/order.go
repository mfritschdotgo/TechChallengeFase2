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

type OrderHandler struct {
	controller *controllers.OrderController
	presenter  interfaces.OrderPresenter
}

func NewOrderHandler(c *controllers.OrderController, p interfaces.OrderPresenter) *OrderHandler {
	return &OrderHandler{
		controller: c,
		presenter:  p,
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var orderDto dto.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&orderDto); err != nil {
		h.presenter.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	order, err := h.controller.CreateOrder(r.Context(), orderDto)
	if err != nil {
		h.presenter.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.presenter.JSON(w, order, http.StatusCreated)
}

func (h *OrderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		h.presenter.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	order, err := h.controller.GetOrderByID(r.Context(), id)
	if err != nil {
		h.presenter.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	h.presenter.JSON(w, order, http.StatusOK)
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil || size <= 0 {
		size = 10
	}

	orders, err := h.controller.GetOrders(r.Context(), page, size)
	if err != nil {
		h.presenter.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	h.presenter.JSON(w, orders, http.StatusOK)
}

func (h *OrderHandler) SetOrderStatus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		h.presenter.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	status, err := strconv.Atoi(chi.URLParam(r, "status"))
	if err != nil {
		h.presenter.Error(w, "Invalid status", http.StatusBadRequest)
		return
	}

	orderStatus, err := h.controller.SetOrderStatus(r.Context(), id, status)
	if err != nil {
		if strings.Contains(err.Error(), "invalid status") {
			h.presenter.Error(w, "Invalid status", http.StatusBadRequest)
		} else {
			h.presenter.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	h.presenter.JSON(w, orderStatus, http.StatusOK)
}

func (h *OrderHandler) FakeCheckout(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		h.presenter.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	_, err := h.controller.SetOrderStatus(r.Context(), id, 4)
	if err != nil {
		h.presenter.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	order, err := h.controller.GetOrderByID(r.Context(), id)
	if err != nil {
		h.presenter.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.presenter.JSON(w, order, http.StatusOK)
}
