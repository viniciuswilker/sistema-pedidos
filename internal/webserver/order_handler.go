package webserver

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/viniciuswilker/sistema-pedidos/internal/dto"
	"github.com/viniciuswilker/sistema-pedidos/internal/entity"
)

type OrderHandler struct {
	OrderDB   entity.OrderRepository
	ProductDB entity.ProductRepository
}

func NewOrderHandler(orderDB entity.OrderRepository, productDB entity.ProductRepository) *OrderHandler {
	return &OrderHandler{OrderDB: orderDB, ProductDB: productDB}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateOrderInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := &entity.Order{
		ID:            uuid.New().String(),
		CustomerCPF:   input.CustomerCPF,
		PaymentMethod: input.PaymentMethod,
		CreatedAt:     time.Now(),
	}

	for _, item := range input.Items {

		product, err := h.ProductDB.FindByID(item.ProductID)
		if err != nil {
			http.Error(w, "Produto não encontrado: "+item.ProductID, http.StatusNotFound)
			return
		}

		orderItem := entity.OrderItem{
			ProductID: product.ID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		}
		order.Items = append(order.Items, orderItem)
		order.Total += product.Price * float64(item.Quantity)
	}

	if err := h.OrderDB.Create(order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {

	orders, err := h.OrderDB.GetAllOrders()
	if err != nil {
		http.Error(w, "Erro ao buscar pedidos", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)

}
