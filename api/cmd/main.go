package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/viniciuswilker/sistema-pedidos/internal/config"
	"github.com/viniciuswilker/sistema-pedidos/internal/database"
	"github.com/viniciuswilker/sistema-pedidos/internal/router"
)

func main() {
	config.CarregarConfigs()
	database.ConectaBanco()

	r := router.CarregarRotas()

	log.Printf("Rodando na porta: %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
