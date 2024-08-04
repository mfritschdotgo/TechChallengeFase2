package presenters

import (
	"encoding/json"
	"net/http"

	"github.com/mfritschdotgo/techchallengefase2/pkg/interfaces"
)

type PaymentPresenter struct{}

func NewPaymentPresenter() interfaces.PaymentPresenter {
	return &PaymentPresenter{}
}

func (p *PaymentPresenter) JSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (p *PaymentPresenter) Error(w http.ResponseWriter, message string, statusCode int) {
	http.Error(w, message, statusCode)
}
