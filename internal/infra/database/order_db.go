package database

import (
	"database/sql"

	"github.com/viniciuswilker/sistema-pedidos/internal/entity"
)

type OrderDB struct {
	DB *sql.DB
}

func NewOrderDB(db *sql.DB) *OrderDB {
	return &OrderDB{DB: db}
}

func (o *OrderDB) Create(order *entity.Order) error {
	tx, err := o.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO orders (id, customer_cpf, total, payment_method, created_at) VALUES (?, ?, ?, ?, ?)",
		order.ID, order.CustomerCPF, order.Total, order.PaymentMethod, order.CreatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range order.Items {
		_, err = tx.Exec("INSERT INTO order_items (order_id, product_id, quantity, price) VALUES (?, ?, ?, ?)",
			order.ID, item.ProductID, item.Quantity, item.Price)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (o *OrderDB) GetByID(id string) (*entity.Order, error) {
	var order entity.Order
	err := o.DB.QueryRow("SELECT id, customer_cpf, total, payment_method, created_at FROM orders WHERE id = ?", id).
		Scan(&order.ID, &order.CustomerCPF, &order.Total, &order.PaymentMethod, &order.CreatedAt)
	if err != nil {
		return nil, err
	}

	rows, err := o.DB.Query("SELECT product_id, quantity, price FROM order_items WHERE order_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item entity.OrderItem
		if err := rows.Scan(&item.ProductID, &item.Quantity, &item.Price); err != nil {
			return nil, err
		}
		order.Items = append(order.Items, item)
	}

	return &order, nil
}
