package postgre

import (
	"database/sql"
	"log"
	"user-service/config"
	"user-service/internal/storage"

	_ "github.com/lib/pq"
)

func NewPostgreRepo(db *sql.DB) *storage.Repository {
	return &storage.Repository{
		UserRepo: NewUserRepository(db),
	}
}

func ConnectPostgre(dbConfig config.DBConfig) *sql.DB {

	db, err := sql.Open(dbConfig.Driver, dbConfig.DSN)
	if err != nil {
		log.Fatalf("DB connection failed, error: %s", err.Error())
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping DB, error: %s", err.Error())
	}
	return db
}
