package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	logger "github.com/blockchaindev100/Go-Blog-Site/logger"
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func RedisInit() {
	Client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"), // Replace with your Redis server address
		Password: "",                         // No password for local development
		DB:       0,                          // Default DB
	})
	ctx := context.Background()
	// Ping the Redis server to check the connection
	pong, err := Client.Ping(ctx).Result()
	if err != nil {
		logger.Logging().Error(err)
		log.Fatal("Error connecting to Redis:", err)
	}
	fmt.Println("Connected to Redis:", pong)
}

func SetData(key string, value string) error {
	err := Client.Set(context.Background(), key, value, time.Hour*24).Err()
	return err
}

func GetData(key string) (string, error) {
	value, err := Client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func RedisConnectionClose() {
	err := Client.Close()
	if err != nil {
		logger.Logging().Error(err)
	}
}
