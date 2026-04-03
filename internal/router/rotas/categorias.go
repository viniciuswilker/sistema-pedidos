package rotas

import (
	"net/http"

	controllers "github.com/viniciuswilker/sistema-pedidos/internal/controllers/api"
)

var rotasCategorias = []Rota{

	{
		URI:                "/categorias",
		Metodo:             []string{http.MethodGet, http.MethodPost},
		Funcao:             controllers.CadastrarCategoria,
		RequerAutenticacao: false,
		TiposPermitidos:    nil,
	},
}
