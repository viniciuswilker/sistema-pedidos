package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/viniciuswilker/sistema-pedidos/internal/database"
	"github.com/viniciuswilker/sistema-pedidos/internal/models"
	"github.com/viniciuswilker/sistema-pedidos/internal/repositorios"
)

func CriarProduto(w http.ResponseWriter, r *http.Request) {

	req, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler a requisição", http.StatusUnprocessableEntity)
		return
	}

	var produto models.Produto
	if err := json.Unmarshal(req, &produto); err != nil {
		http.Error(w, "Erro ao ler a requisição", http.StatusBadRequest)
		return
	}

	if produto.CategoriaID == 0 {
		http.Error(w, "O ID da categoria é obrigatória", http.StatusBadRequest)
		return
	}
	db, err := database.ConectaBanco()
	if err != nil {
		http.Error(w, "Erro de conexão", http.StatusInternalServerError)
		return
	}

	defer db.Close()

	repositorios := repositorios.NovoRepositorioDeProdutos(db)

	id, err := repositorios.Criar(produto)
	if err != nil {
		http.Error(w, "Erro ao criar o produto", http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":       id,
		"mensagem": "Produto criado com sucesso",
	})

}
