package storage

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/behummble/Task-worker/internal/config"
	"github.com/redis/go-redis/v9"
)

type Storage struct {
	log *slog.Logger
	conn *redis.Client
}

func NewStorage(cfg config.RedisConfig, log *slog.Logger) *Storage {
	conn := redis.NewClient(
		&redis.Options{
			Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Username: cfg.Username,
			Password: cfg.Password,
			DialTimeout: time.Duration(cfg.DialTimeout) * time.Second,
			DialerRetries: cfg.Retries,
			ReadTimeout: time.Duration(cfg.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
		},
	)

	return &Storage{
		log: log,
		conn: conn,
	}
}

func (s *Storage) WriteTask([]byte) error {
	
	return nil
}