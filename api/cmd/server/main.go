package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Testando"))
	})

	fmt.Printf("Servidor rodando na porta %s\n", cfg.WebServerPort)
	http.ListenAndServe(":8010", r)
}
