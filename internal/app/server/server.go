package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type Server struct {
	log        *slog.Logger
	httpServer *http.Server
	handler    http.Handler
	Port       int
	Timeout    time.Duration
}

func New(log *slog.Logger, port int, timeout time.Duration, handler http.Handler) *Server {
	return &Server{
		log:     log,
		handler: handler,
		Port:    port,
		Timeout: timeout,
	}
}

func (s *Server) MustRun() {
	err := s.Run()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}

func (s *Server) Run() error {
	const op = "server.Run"
	s.log.With(slog.String("op", op)).
		Info("starting server")

	s.httpServer = &http.Server{
		Addr:           fmt.Sprintf(":%d", s.Port),
		Handler:        s.handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    s.Timeout,
		WriteTimeout:   s.Timeout,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) {
	const op = "server.Shutdown"
	s.log.With(slog.String("op", op)).
		Info("shutting down server")

	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		s.log.Error("failed to shutdown server", slog.String("error", err.Error()))
	}
}
