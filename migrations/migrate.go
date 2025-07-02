package migrations

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"log"
)

func MustRunMigrations(dsn string) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	db.SetMaxOpenConns(1)
	err = goose.SetDialect("postgres")
	if err != nil {
		log.Fatalf("failed to set dialect: %v", err)
	}

	err = goose.Up(db, "migrations")
	if err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	log.Println("Migrations applied successfully")
}
