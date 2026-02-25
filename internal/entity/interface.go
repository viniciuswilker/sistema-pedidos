package entity

type ProductRepository interface {
	FindAll() ([]Product, error)
	FindByCategory(category Category) ([]Product, error)
	FindByID(id string) (*Product, error)
	Create(product *Product) error
}

type OrderRepository interface {
	Create(order *Order) error
	GetByID(id string) (*Order, error)
	GetAllOrders() ([]Order, error)
}
