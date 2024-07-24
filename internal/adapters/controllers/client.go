package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mfritschdotgo/techchallengefase2/internal/adapters/presenters"
	"github.com/mfritschdotgo/techchallengefase2/internal/domain/usecases"
	"github.com/mfritschdotgo/techchallengefase2/pkg/dto"
)

type ClientController struct {
	usecases  *usecases.Client
	presenter presenters.ClientPresenter
}

func NewClientController(s *usecases.Client, p presenters.ClientPresenter) *ClientController {
	return &ClientController{
		usecases:  s,
		presenter: p,
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
func (c *ClientController) CreateClient(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var clientDto dto.CreateClientRequest
	if err := json.NewDecoder(r.Body).Decode(&clientDto); err != nil {
		c.presenter.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	client, err := c.usecases.CreateClient(ctx, clientDto)
	if err != nil {
		c.presenter.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.presenter.JSON(w, client, http.StatusCreated)
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
func (c *ClientController) GetClientByCPF(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cpf := chi.URLParam(r, "cpf")
	if cpf == "" {
		c.presenter.Error(w, "CPF must be provided", http.StatusBadRequest)
		return
	}

	client, err := c.usecases.GetClientByCPF(ctx, cpf)
	if err != nil {
		c.presenter.Error(w, "Client not found", http.StatusNotFound)
		return
	}

	c.presenter.JSON(w, client, http.StatusOK)
}
