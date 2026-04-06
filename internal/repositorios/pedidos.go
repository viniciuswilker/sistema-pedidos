package repositorios

import (
	"database/sql"

	"github.com/viniciuswilker/sistema-pedidos/internal/models"
)

type Pedidos struct {
	db *sql.DB
}

func NovoRepositorioDePedidos(db *sql.DB) *Pedidos {
	return &Pedidos{db}
}

func (repo Pedidos) Criar(pedido models.Pedido) (uint64, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return 0, err
	}

	res, err := tx.Exec(
		"INSERT INTO pedidos (cliente_id, total, status) VALUES (?, ?, ?)",
		pedido.ClienteID, pedido.Total, "pendente",
	)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	pedidoID, _ := res.LastInsertId()

	stmt, err := tx.Prepare("INSERT INTO itens_pedido (pedido_id, produto_id, quantidade, preco_unitario) VALUES (?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	defer stmt.Close()

	for _, item := range pedido.Itens {
		if _, err := stmt.Exec(pedidoID, item.ProdutoID, item.Quantidade, item.PrecoUnitario); err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return uint64(pedidoID), tx.Commit()
}
