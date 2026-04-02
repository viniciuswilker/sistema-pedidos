package models

import "time"

type Produto struct {
	ID          uint32    `json:"id"`
	CategoriaID uint32    `json:"categoria_id"`
	Nome        string    `json:"nome"`
	Descricao   string    `json:"descricao"`
	Preco       float64   `json:"preco"`
	Disponivel  bool      `json:"disponivel"`
	CriadoEm    time.Time `json:"-"`
}
