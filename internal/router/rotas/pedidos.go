package rotas

import (
	"net/http"

	controllers "github.com/viniciuswilker/sistema-pedidos/internal/controllers/api"
)

var rotasPedidos = []Rota{

	{
		URI:                "/pedidos",
		Metodo:             []string{http.MethodPost},
		Funcao:             controllers.CriarPedido,
		RequerAutenticacao: false,
		TiposPermitidos:    nil,
	},
}
