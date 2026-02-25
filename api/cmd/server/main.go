package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/viniciuswilker/sistema-pedidos/configs"
	"github.com/viniciuswilker/sistema-pedidos/internal/infra/database"
	"github.com/viniciuswilker/sistema-pedidos/internal/webserver"
)

func main() {

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

	orderDB := database.NewOrderDB(db)
	orderHandler := webserver.NewOrderHandler(orderDB, productDB)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	fs := http.FileServer(http.Dir("web/static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products", productHandler.GetProducts)

	r.Post("/orders", orderHandler.CreateOrder)
	r.Get("/orders", orderHandler.GetAllOrders)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(
			"web/templates/layout.html",
			"web/templates/index.html",
		)
		if err != nil {
			http.Error(w, "Erro ao carregar templates: "+err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.ExecuteTemplate(w, "layout", nil)
	})

	r.Get("/menu", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles(
			"web/templates/layout.html",
			"web/templates/menu.html",
		))
		tmpl.ExecuteTemplate(w, "layout", nil)
	})

	r.Get("/cart", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(
			"web/templates/layout.html",
			"web/templates/cart.html",
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.ExecuteTemplate(w, "layout", nil)
	})

	r.Get("/orders-list", func(w http.ResponseWriter, r *http.Request) {
		tpml, err := template.ParseFiles(
			"web/templates/layout.html",
			"web/templates/orders.html",
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tpml.ExecuteTemplate(w, "layout", nil)

	})

	fmt.Printf("Servidor rodando na porta %s\n", cfg.WebServerPort)
	http.ListenAndServe(":8010", r)
}
