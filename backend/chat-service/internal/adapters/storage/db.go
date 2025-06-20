package storage

import (
	"chat-service/config"
	"context"
	"crypto/tls"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"time"
)

func ConnectDB(dbConfig config.DBConfig) *sql.DB {
	db, err := sql.Open(dbConfig.Driver, dbConfig.DSN)
	if err != nil {
		panic("Couldn't open database: " + err.Error())
	}
	if err := db.Ping(); err != nil {
		panic("Couldn't ping database: " + err.Error())
	}
	return db
}

func ConnectRedis(redisUrl string) *redis.Client {
	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		panic("Couldn't parse redis url: " + err.Error())
	}
	opt.TLSConfig = &tls.Config{}
	rdb := redis.NewClient(opt)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil {
		panic("Couldn't ping redis: " + err.Error())
	}
	return rdb
}
