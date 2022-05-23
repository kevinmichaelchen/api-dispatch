package db

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
)

type Store struct {
	db          *sql.DB
	redisClient *redis.Client
}

func NewStore(db *sql.DB, redisClient *redis.Client) *Store {
	return &Store{
		db:          db,
		redisClient: redisClient,
	}
}
