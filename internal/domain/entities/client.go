package entities

import (
	"time"
)

type Client struct {
	Name      string    `json:"name" bson:"name"`
	Cpf       string    `json:"cpf" bson:"cpf"`
	Mail      string    `json:"mail" bson:"mail"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func NewClient(name string, cpf string, mail string) (*Client, error) {
	now := time.Now()

	client := &Client{
		Name:      name,
		Cpf:       cpf,
		Mail:      mail,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return client, nil
}
