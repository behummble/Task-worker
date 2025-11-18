package http

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
	"io"

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
	server :=  &Server{
		log: log,
		storage: storage,
	}
	
	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%d", config.Host, config.Port),
		ReadTimeout: time.Duration(config.ReadTimeout) * time.Second,
	}

	mux := newMux(server)
	srv.Handler = mux

	server.server = srv

	return server
}

func(s *Server) Start() {
	s.server.ListenAndServe()
}

func newMux(s *Server) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/task", s.taskHandler)
	return mux
}

func(s *Server) taskHandler(writer http.ResponseWriter, request *http.Request) {
	// TODO:
	//s.ssoClient.Verify(request)
	body, err := request.GetBody()
	defer body.Close()
	if err != nil {
		s.log.Error("Can't execute body from request, client: %s, jobType: %s") //add params
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	data, err := io.ReadAll(body)
	if err != nil {
		s.log.Error("Can't execute bytes from request, client: %s, jobType: %s") //add params
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = s.storage.WriteTask(data)
	if err != nil {
		writer.WriteHeader(http.StatusCreated)
	} else {
		writer.WriteHeader(http.StatusInternalServerError)
	}
}