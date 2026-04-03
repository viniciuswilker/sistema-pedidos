package repositorios

import (
	"database/sql"
)

type Categoria struct {
	db *sql.DB
}

func NovoRepositorioDeCategorias(db *sql.DB) *Categoria {
	return &Categoria{db}
}

func (repositorios Categoria) Criar(nome string) (uint64, error) {

	smtm, err := repositorios.db.Prepare("insert into categorias (nome) values (?)")
	if err != nil {
		return 0, err
	}

	res, err := smtm.Exec(nome)
	if err != nil {
		return 0, err
	}

	categoriaId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(categoriaId), nil
}
