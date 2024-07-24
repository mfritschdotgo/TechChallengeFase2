package presenters

import (
	"encoding/json"
	"net/http"
)

type ClientPresenter struct{}

func NewClientPresenter() ClientPresenter {
	return ClientPresenter{}
}

func (p ClientPresenter) JSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (p ClientPresenter) Error(w http.ResponseWriter, message string, statusCode int) {
	http.Error(w, message, statusCode)
}
