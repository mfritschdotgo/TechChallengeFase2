package presenters

import (
	"encoding/json"
	"net/http"
)

type OrderPresenter struct{}

func NewOrderPresenter() OrderPresenter {
	return OrderPresenter{}
}

func (p OrderPresenter) JSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (p OrderPresenter) Error(w http.ResponseWriter, message string, statusCode int) {
	http.Error(w, message, statusCode)
}
