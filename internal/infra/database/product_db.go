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

func (p *ProductDB) FindAll() ([]entity.Product, error) {
	rows, err := p.DB.Query("SELECT id, name, price, category, description FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Category, &product.Description); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *ProductDB) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.QueryRow("SELECT id, name, price, category, description FROM products WHERE id = ?", id).
		Scan(&product.ID, &product.Name, &product.Price, &product.Category, &product.Description)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductDB) FindByCategory(category entity.Category) ([]entity.Product, error) {
	rows, err := p.DB.Query("SELECT id, name, price, category, description FROM products WHERE category = ?", string(category))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Category, &product.Description); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
