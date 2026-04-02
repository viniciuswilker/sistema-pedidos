package router

import (
	"github.com/gorilla/mux"
	"github.com/viniciuswilker/sistema-pedidos/internal/router/rotas"
)

func CarregarRotas() *mux.Router {

	r := mux.NewRouter()

	r = rotas.Configurar(r)

	return r
}
