package dto

type CreateClientRequest struct {
	Name string `json:"name" bson:"name"`
	Cpf  string `json:"cpf" bson:"cpf"`
	Mail string `json:"mail" bson:"mail"`
}
