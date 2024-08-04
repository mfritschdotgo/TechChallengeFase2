package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mfritschdotgo/techchallengefase2/internal/adapters/controllers"
	"github.com/mfritschdotgo/techchallengefase2/pkg/dto"
	"github.com/mfritschdotgo/techchallengefase2/pkg/interfaces"
)

type PaymentHandler struct {
	controller *controllers.PaymentController
	presenter  interfaces.PaymentPresenter
}

func NewPaymentHandler(c *controllers.PaymentController, p interfaces.PaymentPresenter) *PaymentHandler {
	return &PaymentHandler{
		controller: c,
		presenter:  p,
	}
}

func (h *PaymentHandler) UpdatePaymentStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var paymentDTO dto.PaymentDTO

	if err := json.NewDecoder(r.Body).Decode(&paymentDTO); err != nil {
		h.presenter.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err := h.controller.SetPaymentStatus(ctx, paymentDTO.ID, paymentDTO.Status)
	if err != nil {
		h.presenter.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := map[string]string{"message": "Order payment " + paymentDTO.ID + " updated successfully."}
	h.presenter.JSON(w, response, http.StatusOK)
}

func (h *PaymentHandler) GeneratePayment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		h.presenter.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	qrCode, err := h.controller.GenerateQRCode(r.Context(), id)
	if err != nil {
		h.presenter.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	w.Write(qrCode)
}
