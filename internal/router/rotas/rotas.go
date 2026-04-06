package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI                string
	Metodo             []string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
	TiposPermitidos    []string
}

func Configurar(r *mux.Router) *mux.Router {
	api := r.PathPrefix("/api").Subrouter()
	web := r.PathPrefix("/").Subrouter()

	gruposAPI := [][]Rota{
		rotasUsuarios,
		rotasAuth,
		rotasAdmin,
		rotasCategorias,
		rotasProdutos,
		rotasPedidos,
	}

	for _, grupo := range gruposAPI {
		for _, rota := range grupo {
			var handler http.HandlerFunc = rota.Funcao

			if rota.RequerAutenticacao {
				// handler = middlewares.Autenticar(handler)

				if len(rota.TiposPermitidos) > 0 && rota.TiposPermitidos[0] != "" {
					// handler = middlewares.VerificarPermissao(handler, rota.TiposPermitidos)
				}
			}

			api.HandleFunc(rota.URI, handler).Methods(rota.Metodo...)
		}
	}

	for _, rota := range rotasWeb {
		web.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo...)
	}

	return r
}
