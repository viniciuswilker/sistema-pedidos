package models

import "time"

type ItemPedido struct {
	ID            uint32  `json:"id"`
	PedidoID      uint32  `json:"pedido_id"`
	ProdutoID     uint32  `json:"produto_id"`
	Quantidade    int     `json:"quantidade"`
	PrecoUnitario float64 `json:"preco_unitario"`
}

type Pedido struct {
	ID         uint32       `json:"id"`
	ClienteID  uint32       `json:"cliente_id"`
	DataPedido time.Time    `json:"data_pedido"`
	Status     string       `json:"status"`
	Total      float64      `json:"total"`
	Itens      []ItemPedido `json:"itens"`
}
