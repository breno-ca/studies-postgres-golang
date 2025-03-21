package main

import (
	"log"
	"net/http"

	"studies-postgres-golang/config"
	"studies-postgres-golang/database"
	"studies-postgres-golang/middleware"

	repository "studies-postgres-golang/repository/product"
	service "studies-postgres-golang/service/product"
)

var (
	cfg  = config.NewConfig().WithPostgres()
	pgDB = database.NewPostgres(cfg)

	pr = repository.NewProductRepository(pgDB.DB)
	ps = service.NewProductService(pr)
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", middleware.Logger(http.DefaultServeMux))

	mux.Handle("POST /v1/products", middleware.Logger(http.HandlerFunc(ps.Insert)))
	mux.Handle("GET /v1/products", middleware.Logger(http.HandlerFunc(ps.Index)))

	log.Println("Server running on", cfg.ServerPort)
	if err := http.ListenAndServe(cfg.ServerPort, mux); err != nil {
		log.Println(err.Error())
	}
}
