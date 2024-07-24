package presenters

import (
	"encoding/json"
	"net/http"
)

type CategoryPresenter struct{}

func NewCategoryPresenter() CategoryPresenter {
	return CategoryPresenter{}
}

func (p CategoryPresenter) JSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (p CategoryPresenter) Error(w http.ResponseWriter, message string, statusCode int) {
	http.Error(w, message, statusCode)
}
