package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/viniciuswilker/sistema-pedidos/internal/database"
	"github.com/viniciuswilker/sistema-pedidos/internal/models"
	"github.com/viniciuswilker/sistema-pedidos/internal/repositorios"
)

func CriarPedido(w http.ResponseWriter, r *http.Request) {

	req, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler corpo", http.StatusBadRequest)
		return
	}

	var pedido models.Pedido
	if err := json.Unmarshal(req, &pedido); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	db, err := database.ConectaBanco()
	if err != nil {
		http.Error(w, "Erro no banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	repositorios := repositorios.NovoRepositorioDePedidos(db)
	id, err := repositorios.Criar(pedido)
	if err != nil {
		fmt.Println("Erro real no banco:", err)
		http.Error(w, "Erro ao processar pedido", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]uint64{"id": id})
}
