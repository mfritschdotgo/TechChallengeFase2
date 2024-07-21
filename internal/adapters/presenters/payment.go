package presenters

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mfritschdotgo/techchallengefase2/internal/adapters/dto"
	"github.com/mfritschdotgo/techchallengefase2/internal/domain/usecases"
)

type PaymentHandler struct {
	usecases *usecases.Payment
}

func NewPaymentHandler(s *usecases.Payment) *PaymentHandler {
	return &PaymentHandler{
		usecases: s,
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
func (h *PaymentHandler) UpdatePaymentStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var paymentDTO dto.PaymentDTO

	if err := json.NewDecoder(r.Body).Decode(&paymentDTO); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err := h.usecases.SetPaymentStatus(ctx, paymentDTO.ID, paymentDTO.Status)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": "Order payment " + paymentDTO.ID + " updated successfully."}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
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
func (h *PaymentHandler) GeneratePayment(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}
	qrCode, err := h.usecases.GenerateQRCode(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	w.Write(qrCode)
}
