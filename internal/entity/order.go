package entity

import "time"

type OrderItem struct {
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

type Order struct {
	ID            string      `json:"id"`
	CustomerCPF   string      `json:"customer_cpf"`
	CustomerName  string      `json:"customer_name"`
	Items         []OrderItem `json:"items"`
	Total         float64     `json:"total"`
	PaymentMethod string      `json:"payment_method"`
	CreatedAt     time.Time   `json:"created_at"`
}
