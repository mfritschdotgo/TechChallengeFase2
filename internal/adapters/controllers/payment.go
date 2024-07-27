package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mfritschdotgo/techchallengefase2/internal/adapters/presenters"
	"github.com/mfritschdotgo/techchallengefase2/internal/domain/usecases"
	"github.com/mfritschdotgo/techchallengefase2/pkg/dto"
)

type PaymentController struct {
	usecases  *usecases.Payment
	presenter presenters.PaymentPresenter
}

func NewPaymentController(s *usecases.Payment, p presenters.PaymentPresenter) *PaymentController {
	return &PaymentController{
		usecases:  s,
		presenter: p,
	}
}

// UpdatePaymentStatus Update payment status for the order.
// @Summary Update payment status for the order.
// @Description Update payment status for the order based ID.
// @Tags payment
// @Accept json
// @Produce json
// @Param		request	body		dto.PaymentDTO	true	"New payment status for the order"
// @Success 200 "Successfully update payment status for the order"
// @Failure 400 "Bad request if the ID is not provided or invalid"
// @Failure 500 "Internal server error if there is a problem on the server side"
// @Router /payment [post]
func (c *PaymentController) UpdatePaymentStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var paymentDTO dto.PaymentDTO

	if err := json.NewDecoder(r.Body).Decode(&paymentDTO); err != nil {
		c.presenter.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err := c.usecases.SetPaymentStatus(ctx, paymentDTO.ID, paymentDTO.Status)
	if err != nil {
		c.presenter.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := map[string]string{"message": "Order payment " + paymentDTO.ID + " updated successfully."}
	c.presenter.JSON(w, response, http.StatusOK)
}

// generatePayment Generates the qr code for payment via pix
// @Summary Generates the qr code for payment via pix
// @Description Generates the qr code for payment via pix
// @Tags payment
// @Accept json
// @Produce json
// @Param id path string true "order ID"
// @Success 200 "Got qr code successfully"
// @Failure 400 "Bad request if the ID is not provided or invalid"
// @Failure 500 "Internal server error if there is a problem on the server side"
// @Router /payment/{id} [get]
func (c *PaymentController) GeneratePayment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		c.presenter.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	qrCode, err := c.usecases.GenerateQRCode(r.Context(), id)
	if err != nil {
		c.presenter.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	w.Write(qrCode)
}
