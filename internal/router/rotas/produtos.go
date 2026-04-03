package rotas

import (
	"net/http"

	controllers "github.com/viniciuswilker/sistema-pedidos/internal/controllers/api"
)

var rotasProdutos = []Rota{

	{
		URI:                "/produtos",
		Metodo:             []string{http.MethodPost},
		Funcao:             controllers.CriarProduto,
		RequerAutenticacao: false,
		TiposPermitidos:    nil,
	},

	{
		URI:                "/produtos",
		Metodo:             []string{http.MethodGet},
		Funcao:             func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("GET EM PRODUTOS")) },
		RequerAutenticacao: false,
		TiposPermitidos:    nil,
	},


	{
		URI:                "/produtos/{categoriaID}",
		Metodo:             []string{http.MethodGet},
		Funcao:             func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("GET EM PRODUTOS")) },
		RequerAutenticacao: false,
		TiposPermitidos:    nil,
	},
}
