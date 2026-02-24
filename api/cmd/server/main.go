package main

import (
	"database/sql"
	"fmt"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gofiber/fiber/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/viniciuswilker/sistema-pedidos/configs"
	"github.com/viniciuswilker/sistema-pedidos/internal/infra/database"
	"github.com/viniciuswilker/sistema-pedidos/internal/webserver"
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

	productDB := database.NewProductDB(db)
	productHandler := webserver.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/products", productHandler.CreateProduct)

	fmt.Printf("Servidor rodando na porta %s\n", cfg.WebServerPort)
}
