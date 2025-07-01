package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func InitDB(dsn string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
