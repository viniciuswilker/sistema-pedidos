package database

import (
	"database/sql"

	"github.com/viniciuswilker/sistema-pedidos/internal/entity"
)

type ProductDB struct {
	DB *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{DB: db}
}

func (p *ProductDB) Create(product *entity.Product) error {
	stmt, err := p.DB.Prepare("INSERT INTO products (id, name, price, category, description) VALUES (?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price, product.Category, product.Description)
	return err

}

func (p *ProductDB) FindByCategory(category string) ([]entity.Product, error) {
	rows, err := p.DB.Query("SELECT id, name, price, category, description from product where category = ?", category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entity.Product

	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
