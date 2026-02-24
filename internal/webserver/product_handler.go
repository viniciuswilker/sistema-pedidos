package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/viniciuswilker/sistema-pedidos/internal/entity"
)

type ProductHandler struct {
	ProductDB entity.ProductRepository
}

func NewProductHandler(db entity.ProductRepository) *ProductHandler {
	return &ProductHandler{ProductDB: db}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	product.ID = uuid.New().String()
	if err := h.ProductDB.Create(&product); err != nil {
		http.Error(w, "Erro ao salvar produto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	var products []entity.Product
	var err error

	if category != "" {
		products, err = h.ProductDB.FindByCategory(entity.Category(category))
	} else {
		products, err = h.ProductDB.FindAll()
	}

	if err != nil {
		http.Error(w, "Erro ao buscar produtos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
