package entity

type ProductRepository interface {
	FindAll() ([]Product, error)
	FindByCategory(category Category) ([]Product, error)
	FindByID(id string) (*Product, error)
}

type OrderRepository interface {
	Create(order *Order) error
	GetBtID(id string) (*Order, error)
}
