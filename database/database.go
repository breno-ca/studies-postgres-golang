package database

import (
	"database/sql"
	"log"

	"studies-postgres-golang/config"

	_ "github.com/lib/pq"
)

type postgres struct {
	DB *sql.DB
}

func NewPostgres(config config.Config) *postgres {
	db, err := sql.Open("postgres", config.PostgresConfig.DSN())
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %s", err)
	}

	return &postgres{
		DB: db,
	}
}
