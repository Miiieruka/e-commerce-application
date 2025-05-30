package storage

import (
	"context"
	"crypto/tls"
	"database/sql"
	"log"
	"product-service/config"
	"product-service/internal/storage"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRepository(db *sql.DB, rdb *redis.Client) *storage.Repository {
	return &storage.Repository{
		ProductRepo: NewProductRepository(db, rdb),
	}
}

func ConnectPostgre(dbConfig config.DBConfig) *sql.DB {
	db, err := sql.Open(dbConfig.Driver, dbConfig.DSN)

	if err != nil {
		log.Fatalf("ConnectPostgre: %s", err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("ConnectPostgrePing: %s", err.Error())
	}
	return db
}

func ConnectRedis(redisUrl string) *redis.Client {
	opt, err := redis.ParseURL(redisUrl)

	if err != nil {
		log.Fatalf("ConnectRedis: %s", err.Error())
	}

	opt.TLSConfig = &tls.Config{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rdb := redis.NewClient(opt)
	if err = rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("ConnectRedisPing: %s", err.Error())
	}

	return rdb
}
