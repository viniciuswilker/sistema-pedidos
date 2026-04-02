package router

import "github.com/gorilla/mux"

func CarregarRotas() *mux.Router {

	r := mux.NewRouter()

	r := rotas.Configurar()

	return r
}
