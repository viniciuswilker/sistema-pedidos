package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/viniciuswilker/sistema-pedidos/internal/database"
	"github.com/viniciuswilker/sistema-pedidos/internal/models"
	"github.com/viniciuswilker/sistema-pedidos/internal/repositorios"
)

func CadastrarCategoria(w http.ResponseWriter, r *http.Request) {

	req, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler corpo da requisição", http.StatusBadRequest)
		return
	}

	var categoria models.Categoria
	if err := json.Unmarshal(req, &categoria); err != nil {
		http.Error(w, "Erro ao ler corpo da requisição", http.StatusBadRequest)
		return
	}

	db, err := database.ConectaBanco()
	if err != nil {
		http.Error(w, "Erro ao ler corpo da requisição", http.StatusBadRequest)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeCategorias(db)

	categoriaID, err := repositorio.Criar(categoria.Nome)
	if err != nil {
		http.Error(w, "Erro ao ler corpo da requisição", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"mensagem":     "Categoria cadastrada com sucesso",
		"categoria_id": categoriaID,
	})
}
