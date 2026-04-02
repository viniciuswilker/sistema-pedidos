package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/viniciuswilker/sistema-pedidos/internal/config"
	"github.com/viniciuswilker/sistema-pedidos/internal/database"
)

func main() {
	config.CarregarConfigs()
	database.ConectaBanco()

	r := mux.NewRouter()

	log.Printf("Rodando na porta: %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
