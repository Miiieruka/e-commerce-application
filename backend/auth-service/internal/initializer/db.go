package initializer

import (
	confid "auth-service/internal/config"
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"
)

var (
	ErrDBConnection = errors.New("failed connecting to db")
)

func ConnectDB(config confid.Config) (*sql.DB, error) {
	dbConfig := config.DB
	db, err := sql.Open(dbConfig.Driver, dbConfig.DSN)
	if err != nil {
		log.Fatalf("DB connection failed, error: %s", err.Error())
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping DB, error: %s", err.Error())
	}
	return db, nil
}

func CloseConnection(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Fatalf("Error closing DB connection")
	}
}
