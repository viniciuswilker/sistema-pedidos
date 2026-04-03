package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func ListarCategorias(w http.ResponseWriter, r *http.Request) {

	db, err := database.ConectaBanco()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeCategorias(db)

	categorias, err := repositorio.Listar()
	if err != nil {
		http.Error(w, "Erro ao listar categorias", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(categorias); err != nil {
		http.Error(w, "Erro ao converter resposta para JSON", http.StatusInternalServerError)
		return
	}
}

func DeletarCategoria(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	categoriaID, err := strconv.ParseUint(parametros["categoriaID"], 10, 64)
	if err != nil {
		http.Error(w, "Erro ao converter parametro da URL", http.StatusBadRequest)
		return
	}

	db, err := database.ConectaBanco()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	repositorios := repositorios.NovoRepositorioDeCategorias(db)

	if err := repositorios.Deletar(categoriaID); err != nil {
		http.Error(w, "Erro ao deletar a categoria", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
