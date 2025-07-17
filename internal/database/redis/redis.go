package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"service-advert/internal/config"
	"service-advert/internal/models"
	"strconv"
	"time"

	//models "service-advert/internal/models"
	redis "github.com/redis/go-redis/v9"
)

var (
	db  *redis.Client
	ttl int
)

func Init(cfg *config.Config) error {
	ttl = cfg.RDS.TTL
	db = redis.NewClient(&redis.Options{
		Addr: cfg.RDS.Addr,
		// Password:     cfg.RDS.Password,
		DB: cfg.RDS.DB,
		// Username:     cfg.RDS.User,
		MaxRetries:   cfg.RDS.MaxRetries,
		DialTimeout:  time.Duration(cfg.RDS.DialTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.RDS.Timeout) * time.Second,
		WriteTimeout: time.Duration(cfg.RDS.Timeout) * time.Second,
	})

	if err := db.Ping(context.Background()).Err(); err != nil {
		return fmt.Errorf("failed to connect to redis server: %s", err.Error())
	}
	slog.Info("Connect to Redis")
	return nil
}

func CloseDB() {
	db.Close()
}

func Set(advert *models.Advert) error {
	// Переводим структуру в json
	slog.Info("Redis: Set", "good", advert.ID)
	value, err := json.Marshal(&advert)
	if err != nil {
		return err
	}

	key := strconv.Itoa(advert.ID)
	if err := db.Set(context.Background(), key, value, time.Duration(ttl)*time.Second).Err(); err != nil {
		slog.Error("Redis Set: failed to set data, error:", "err", err.Error())
		return fmt.Errorf("redis: failed to set data, error: %v", err)
	}
	return nil
}

func Get(id string) (models.Advert, error) {
	val, err := db.Get(context.Background(), id).Result()

	if err == redis.Nil {
		slog.Info("Redis Get: value not found")
		return models.Advert{}, fmt.Errorf("redis: value not found")
	} else if err != nil {
		slog.Error("Redis Get: failed to get value, error:", "err", err.Error())
		return models.Advert{}, fmt.Errorf("redis: failed to get value, error: %v", err)
	}
	// Переводим из json в структуру modeldb.Goods
	var advert models.Advert
	if err = json.Unmarshal([]byte(val), &advert); err != nil {
		slog.Error(err.Error())
		return models.Advert{}, err
	}

	slog.Info("Redis Get", "id", id)
	return advert, nil
}
