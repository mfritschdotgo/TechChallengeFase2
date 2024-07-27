package presenters

import (
	"encoding/json"
	"net/http"
)

type ProductPresenter struct{}

func NewProductPresenter() ProductPresenter {
	return ProductPresenter{}
}

func (p ProductPresenter) JSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (p ProductPresenter) Error(w http.ResponseWriter, message string, statusCode int) {
	http.Error(w, message, statusCode)
}
