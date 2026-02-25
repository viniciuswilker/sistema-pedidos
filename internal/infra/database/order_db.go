package database

import (
	"database/sql"
	"fmt"

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

	_, err = tx.Exec("INSERT INTO orders (id, customer_cpf, customer_name ,total, payment_method, order_status ,created_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
		order.ID, order.CustomerCPF, order.CustomerName, order.Total, order.PaymentMethod, "PENDENTE", order.CreatedAt)
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
	err := o.DB.QueryRow("SELECT id, customer_cpf, customer_name, total, payment_method, order_status,created_at FROM orders WHERE id = ?", id).
		Scan(&order.ID, &order.CustomerCPF, &order.CustomerName, &order.Total, &order.PaymentMethod, &order.OrderStatus, &order.CreatedAt)
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

func (o *OrderDB) GetAllOrders() ([]entity.Order, error) {

	rows, err := o.DB.Query("SELECT id, customer_cpf, customer_name, total, payment_method, order_status, created_at FROM orders")
	if err != nil {
		fmt.Printf("Erro no Scan do pedido: %v\n", err)

		return nil, err
	}
	defer rows.Close()

	var orders []entity.Order
	for rows.Next() {
		var order entity.Order

		err := rows.Scan(&order.ID, &order.CustomerCPF, &order.CustomerName, &order.Total, &order.PaymentMethod, &order.OrderStatus, &order.CreatedAt)
		if err != nil {
			return nil, err
		}

		itemRows, err := o.DB.Query(`
            SELECT i.product_id, p.name, i.quantity, i.price 
            FROM order_items i 
            INNER JOIN products p ON i.product_id = p.id 
            WHERE i.order_id = ?`, order.ID)
		if err != nil {
			fmt.Printf("Erro no Scan do pedido: %v\n", err)
			return nil, err
		}

		for itemRows.Next() {
			var item entity.OrderItem
			if err := itemRows.Scan(&item.ProductID, &item.ProductName, &item.Quantity, &item.Price); err != nil {
				itemRows.Close()
				fmt.Printf("Erro no Scan do pedido: %v\n", err)

				return nil, err
			}
			order.Items = append(order.Items, item)
		}
		itemRows.Close()

		orders = append(orders, order)
	}

	return orders, nil
}
