package worker

import (
	"log/slog"
)

func NewWorkerPool(log *slog.Logger) *WorkerPool {
	return &WorkerPool{}
}

type WorkerPool struct {
	log *slog.Logger
	workerCount int
}