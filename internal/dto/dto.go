package dto

type OrderItemInput struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type CreateOrderInput struct {
	CustomerCPF   string           `json:"customer_cpf"`
	Items         []OrderItemInput `json:"items"`
	PaymentMethod string           `json:"payment_method"`
}
