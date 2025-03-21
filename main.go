package main

import (
	"log"
	"net/http"

	"studies-postgres-golang/config"
	"studies-postgres-golang/database"

	repository "studies-postgres-golang/repository/product"
	service "studies-postgres-golang/service/product"
)

var (
	cfg  = config.NewConfig().WithPostgres()
	pgDB = database.NewPostgres(cfg)

	productRepository = repository.NewProductRepository(pgDB.DB)
	productService    = service.NewProductService(productRepository)
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.DefaultServeMux)

	mux.Handle("POST /v1/products", http.HandlerFunc(productService.Insert))
	mux.Handle("GET /v1/products", http.HandlerFunc(productService.Index))

	log.Println("Server running on", cfg.ServerPort)
	if err := http.ListenAndServe(cfg.ServerPort, mux); err != nil {
		log.Println(err.Error())
	}
}
