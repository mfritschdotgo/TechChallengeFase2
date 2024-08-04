package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mfritschdotgo/techchallengefase2/internal/adapters/controllers"
	"github.com/mfritschdotgo/techchallengefase2/internal/adapters/presenters"
	"github.com/mfritschdotgo/techchallengefase2/pkg/dto"
)

type ClientHandler struct {
	controller *controllers.ClientController
	presenter  presenters.ClientPresenter
}

func NewClientHandler(c *controllers.ClientController, p presenters.ClientPresenter) *ClientHandler {
	return &ClientHandler{
		controller: c,
		presenter:  p,
	}
}

func (h *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var clientDto dto.CreateClientRequest
	if err := json.NewDecoder(r.Body).Decode(&clientDto); err != nil {
		h.presenter.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	client, err := h.controller.CreateClient(r.Context(), clientDto)
	if err != nil {
		h.presenter.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.presenter.JSON(w, client, http.StatusCreated)
}

func (h *ClientHandler) GetClientByCPF(w http.ResponseWriter, r *http.Request) {
	cpf := chi.URLParam(r, "cpf")
	if cpf == "" {
		h.presenter.Error(w, "CPF must be provided", http.StatusBadRequest)
		return
	}

	client, err := h.controller.GetClientByCPF(r.Context(), cpf)
	if err != nil {
		h.presenter.Error(w, "Client not found", http.StatusNotFound)
		return
	}

	h.presenter.JSON(w, client, http.StatusOK)
}
