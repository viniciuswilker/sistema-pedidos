package rotas

import (
	"net/http"

	controllers "github.com/viniciuswilker/sistema-pedidos/internal/controllers/api"
)

var rotasCategorias = []Rota{

	{
		URI:                "/categorias",
		Metodo:             []string{http.MethodGet},
		Funcao:             controllers.ListarCategorias,
		RequerAutenticacao: false,
		TiposPermitidos:    nil,
	},

	{
		URI:                "/categorias",
		Metodo:             []string{http.MethodPost},
		Funcao:             controllers.CadastrarCategoria,
		RequerAutenticacao: false,
		TiposPermitidos:    nil,
	},

	{
		URI:                "/categorias/{categoriaID}",
		Metodo:             []string{http.MethodDelete},
		Funcao:             controllers.DeletarCategoria,
		RequerAutenticacao: false,
		TiposPermitidos:    nil,
	},
}
