package repositorios

import (
	"database/sql"

	"github.com/viniciuswilker/sistema-pedidos/internal/models"
)

type produtos struct {
	db *sql.DB
}

func NovoRepositorioDeProdutos(db *sql.DB) *produtos {
	return &produtos{db}
}

func (repositorio produtos) Criar(produto models.Produto) (uint64, error) {

	stmt, err := repositorio.db.Prepare("insert into produtos (categoria_id, nome, descricao, preco, disponivel) values (?,?,?,?,?)")
	if err != nil {
		return 0, nil
	}
	defer stmt.Close()

	res, err := stmt.Exec(produto.CategoriaID, produto.Nome, produto.Descricao, produto.Preco, produto.Disponivel)
	if err != nil {
		return 0, nil
	}

	pedidoId, err := res.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return uint64(pedidoId), nil

}
