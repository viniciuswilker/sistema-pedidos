package repositorios

import (
	"database/sql"

	"github.com/viniciuswilker/sistema-pedidos/internal/models"
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

func (repositorios Categoria) Deletar(categoriaID uint64) error {

	smtm, err := repositorios.db.Prepare("delete from categorias where id = ?")
	if err != nil {
		return err
	}

	if _, err := smtm.Exec(categoriaID); err != nil {
		return err
	}

	return nil
}

func (repositorios Categoria) Listar() ([]models.Categoria, error) {

	rows, err := repositorios.db.Query("select id, nome from categorias")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categorias []models.Categoria

	for rows.Next() {
		var categoria models.Categoria

		if err := rows.Scan(&categoria.ID, &categoria.Nome); err != nil {
			return nil, err
		}

		categorias = append(categorias, categoria)

	}

	return categorias, nil

}
