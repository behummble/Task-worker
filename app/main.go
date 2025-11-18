package main

import (
	"log/slog"
	"os"

	"github.com/behummble/Task-worker/internal/handlers/http"
	"github.com/behummble/Task-worker/internal/config"
	"github.com/behummble/Task-worker/internal/storage"
)

func main() {
	cfg := config.NewConfig("")
	logger := newLogger(cfg.Log)
	storage := storage.NewStorage(cfg.Redis, logger)
	server := http.NewServer(cfg.Server, logger, storage)
	server.Start()
	
	//jobRegister := newJobRegister()
	//worker := NewWorkerPool()
	//worker.Start()
}

func newLogger(cfg config.LogConfig) *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{Level: slog.Level(cfg.LogLevel)},
		),
	)
}

type Job interface {
	Process() error
}
