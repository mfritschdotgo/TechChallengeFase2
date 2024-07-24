package dto

type ProductItem struct {
	ID          string `json:"id"`
	Quantity    int64  `json:"quantity"`
	ProductName string `json:"product_name" swaggerignore:"true"`
}

type CreateOrderRequest struct {
	Client   string        `json:"client"`
	Products []ProductItem `json:"products"`
}
