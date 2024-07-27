package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/mfritschdotgo/techchallengefase2/internal/adapters/presenters"
	"github.com/mfritschdotgo/techchallengefase2/internal/domain/usecases"
	"github.com/mfritschdotgo/techchallengefase2/pkg/dto"
)

type OrderController struct {
	usecases  *usecases.Order
	presenter presenters.OrderPresenter
}

func NewOrderController(s *usecases.Order, p presenters.OrderPresenter) *OrderController {
	return &OrderController{
		usecases:  s,
		presenter: p,
	}
}

// CreateOrder adds a new order to the store
// @Summary Add a new order
// @Description Adds a new order to the database with the given details.
// @Tags orders
// @Accept json
// @Produce json
// @Param		request	body		dto.CreateOrderRequest	true	"Order creation details"
// @Success 201 {object} entities.Order "Successfully created Order"
// @Failure 400 "Bad request if the Order data is invalid"
// @Failure 500 "Internal server error if there is a problem on the server side"
// @Router /orders [post]
func (c *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var orderDto dto.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&orderDto); err != nil {
		c.presenter.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	order, err := c.usecases.CreateOrder(ctx, orderDto)
	if err != nil {
		c.presenter.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.presenter.JSON(w, order, http.StatusCreated)
}

// GetOrderByID retrieves a order by its ID
// @Summary Get a order
// @Description Retrieves details of a order based on its unique ID.
// @Tags orders
// @Accept json
// @Produce json
// @Param id path string true "order ID"
// @Success 200 {object} entities.Order "Successfully retrieved the order details"
// @Failure 400 "Bad request if the ID is not provided or invalid"
// @Failure 404 "Product not found if the ID does not match any order"
// @Router /orders/{id} [get]
func (c *OrderController) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	if id == "" {
		c.presenter.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	order, err := c.usecases.GetOrderByID(ctx, id)
	if err != nil {
		c.presenter.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	c.presenter.JSON(w, order, http.StatusOK)
}

// GetOrders retrieves a list of orders
// @Summary List orders
// @Description Retrieves a paginated list of orders
// @Tags orders
// @Accept json
// @Produce json
// @Param page query int false "Page number for pagination" default(1)
// @Param pageSize query int false "Number of orders per page" default(10)
// @Success 200 {array} entities.Order "Successfully retrieved list of orders"
// @Failure 500 "Internal server error if there is a problem on the server side"
// @Router /orders [get]
func (c *OrderController) GetOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil || size <= 0 {
		size = 10
	}

	orders, err := c.usecases.GetOrders(ctx, page, size)
	if err != nil {
		c.presenter.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	c.presenter.JSON(w, orders, http.StatusOK)
}

// update order status
// @Summary Update order status
// @Description Update order status, statuses 1 to 4 allowed
// @Tags orders
// @Accept json
// @Produce json
// @Param id path string true "order ID"
// @Param status path string true "order ID"
// @Success 200 {object} entities.OrderStatus "Successfully status updated"
// @Failure 400 "Bad request if the ID is not provided or invalid"
// @Failure 400 "Bad request if the Status is not provided or invalid"
// @Failure 500 "Internal server error if there is a problem on the server side"
// @Router /orders/{id}/{status} [patch]
func (c *OrderController) SetOrderStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	if id == "" {
		c.presenter.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	status, err := strconv.Atoi(chi.URLParam(r, "status"))
	if err != nil {
		c.presenter.Error(w, "Invalid status", http.StatusBadRequest)
		return
	}

	orderStatus, err := c.usecases.SetOrderStatus(ctx, id, status)

	if err != nil {
		if strings.Contains(err.Error(), "invalid status") {
			c.presenter.Error(w, "Invalid status", http.StatusBadRequest)
		} else {
			c.presenter.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	c.presenter.JSON(w, orderStatus, http.StatusOK)
}

// simulates a checkout
// @Summary Simulates a checkout
// @Description Simulates a checkout, changing status to 4 - finished
// @Tags fakeCheckout
// @Accept json
// @Produce json
// @Param id path string true "order ID"
// @Success 200 {object} entities.Order "Successfully fake checkout"
// @Failure 400 "Bad request if the ID is not provided or invalid"
// @Failure 500 "Internal server error if there is a problem on the server side"
// @Router /fakeCheckout/{id} [post]
func (c *OrderController) FakeCheckout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	if id == "" {
		c.presenter.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	_, err := c.usecases.SetOrderStatus(ctx, id, 1)

	if err != nil {
		c.presenter.Error(w, err.Error(), http.StatusBadRequest)
	}

	order, err := c.usecases.GetOrderByID(ctx, id)

	if err != nil {
		c.presenter.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.presenter.JSON(w, order, http.StatusOK)
}
