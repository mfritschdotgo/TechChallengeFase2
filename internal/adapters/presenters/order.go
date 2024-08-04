package presenters

import (
	"encoding/json"
	"net/http"

	"github.com/mfritschdotgo/techchallengefase2/pkg/interfaces"
)

type OrderPresenter struct{}

func NewOrderPresenter() interfaces.OrderPresenter {
	return &OrderPresenter{}
}

func (p *OrderPresenter) JSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (p *OrderPresenter) Error(w http.ResponseWriter, message string, statusCode int) {
	http.Error(w, message, statusCode)
}
