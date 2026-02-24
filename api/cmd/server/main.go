package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/viniciuswilker/sistema-pedidos/configs"
)

func main() {
	fmt.Printf("Servidor INICIANDO\n")

	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("sqlite3", cfg.DBName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Printf("Servidor rodando na porta %s\n", cfg.WebServerPort)
}
