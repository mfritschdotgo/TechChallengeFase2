package presenters

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/mfritschdotgo/techchallengefase2/internal/adapters/dto"
	"github.com/mfritschdotgo/techchallengefase2/internal/domain/usecases"
)

type ClientHandler struct {
	usecases *usecases.Client
}

func NewClientHandler(s *usecases.Client) *ClientHandler {
	return &ClientHandler{
		usecases: s,
	}
}

// CreateClient adds a new client to the store
// @Summary Add a new client
// @Description Adds a new client to the database with the given details.
// @Tags clients
// @Accept json
// @Produce json
// @Param		request	body		dto.CreateClientRequest	true	"Client creation details"
// @Success 201 {object} entities.Client "Client successfully created"
// @Failure 400 "Bad request if the Client data is invalid"
// @Failure 500 "Internal server error if there is a problem on the server side"
// @Router /clients [post]
func (h *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // Get the request context

	var clientDto dto.CreateClientRequest
	if err := json.NewDecoder(r.Body).Decode(&clientDto); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	client, err := h.usecases.CreateClient(ctx, clientDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(client)
}

// GetClientByCPF retrieves a client by its CPF
// @Summary Get a client
// @Description Retrieves details of a client based on its unique CPF.
// @Tags clients
// @Accept json
// @Produce json
// @Param cpf path string true "client CPF"
// @Success 200 {object} entities.Client "Successfully retrieved the client details"
// @Failure 400 "Bad request if the CPF is not provided or invalid"
// @Failure 404 "Client not found if the CPF does not match any Client"
// @Router /clients/{cpf} [get]
func (h *ClientHandler) GetClientByCPF(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cpf := chi.URLParam(r, "cpf")
	if cpf == "" {
		http.Error(w, "CPF must be provided", http.StatusBadRequest)
		return
	}

	client, err := h.usecases.GetClientByCPF(ctx, cpf)
	if err != nil {
		http.Error(w, "Client not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(client)
}
