package redis

import (
	"Amooz/pkg/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"time"
)

type RedisDb struct {
	client *redis.Client
}

func NewRedis(config config.RedisConfig) RedisDb {
	dbName, _ := strconv.Atoi(config.DbName)
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.Host, config.Port),
		//Username:     config.Username,
		Password:     config.Pass,
		DB:           dbName,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}
	fmt.Println("Connected to Redis:", pong)

	return RedisDb{client: rdb}
}

func (a RedisDb) Client() *redis.Client {
	return a.client
}
