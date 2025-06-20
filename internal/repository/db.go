package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // драйвер для PostgreSQL
	"log"
)

func InitDB(dsn string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	
	return db
}
