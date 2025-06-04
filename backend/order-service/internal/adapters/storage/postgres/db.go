package postgres

import (
	"database/sql"
	"log"
	"order-service/config"

	_ "github.com/lib/pq"
)

func ConnectPostgres(dbConfig config.DBConfig) *sql.DB {
	db, err := sql.Open(dbConfig.Driver, dbConfig.DSN)
	if err != nil {
		log.Fatalf("Couldn't open database")
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Couldn't ping db")
	}
	return db
}
