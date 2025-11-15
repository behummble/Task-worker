package http

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/behummble/Task-worker/internal/config"
)

type Server struct {
	server *http.Server
	storage Storage
	log *slog.Logger
}

type Storage interface {
	WriteTask([]byte) error
}

func NewServer(config config.ServerConfig, log *slog.Logger, storage Storage) *Server {
	mux := newMux()
	
	srv := &http.Server{
		Handler: mux,
		Addr: fmt.Sprintf("%s:%d", config.Host, config.Port),
		ReadTimeout: time.Duration(config.ReadTimeout) * time.Second,
	}

	return &Server{
		server: srv,
		log: log,
		storage: storage,
	}
}

func(s *Server) Start() {
	s.server.ListenAndServe()
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/task", taskHandler)
	return mux
}

func taskHandler(writer http.ResponseWriter, request *http.Request) {
	
}