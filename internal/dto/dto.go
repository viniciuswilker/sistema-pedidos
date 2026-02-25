package dto

type OrderItemInput struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type CreateOrderInput struct {
	CustomerCPF   string           `json:"customer_cpf"`
	CustomerName  string           `json:"customer_name"`
	Items         []OrderItemInput `json:"items"`
	PaymentMethod string           `json:"payment_method"`
}
