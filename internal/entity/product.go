package entity

type Category string

const (
	Food     Category = "food"
	Drink    Category = "drink"
	Addition Category = "addition"
)

type Product struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Category    Category `json:"category"`
	Description string   `json:"description"`
}
