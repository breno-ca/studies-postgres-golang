package main

import (
	"studies-postgres-golang/config"
	"studies-postgres-golang/database"
)

var (
	cfg  = config.NewConfig().WithPostgres()
	pgDB = database.NewPostgres(cfg)
)

func main() {
}
