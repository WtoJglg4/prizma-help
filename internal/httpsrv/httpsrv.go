package httpsrv

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

type Server struct {
	router *http.ServeMux
	server *http.Server
	logger *slog.Logger
}

type Options struct {
	Addr   string
	Logger *slog.Logger
}

func New(opts Options) *Server {
	mux := http.NewServeMux()
	//nolint:gosec // Slowloris attack probably should not be mitigated, because we are not exposing this service?
	srv := &http.Server{
		Handler: mux,
		Addr:    opts.Addr,
	}
	return &Server{
		router: mux,
		server: srv,
		logger: opts.Logger.With("addr", opts.Addr),
	}
}

func (s *Server) Register(
	path string,
	handler http.Handler,
	middlewares ...func(handler http.Handler) http.Handler,
) {
	for _, m := range middlewares {
		handler = m(handler)
	}
	s.router.Handle(path, handler)
}

func (s *Server) RegisterFunc(
	path string,
	handler http.HandlerFunc,
	middlewares ...func(handlerFunc http.HandlerFunc) http.HandlerFunc,
) {
	for _, m := range middlewares {
		handler = m(handler)
	}
	s.router.Handle(path, handler)
}

func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("starting HTTP server")
	go func() {
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error("http server ListenAndServe failed", "err", err)
		}
	}()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Info("stopping HTTP server on")
	err := s.server.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("http server Shutdown failed: %w", err)
	}
	return nil
}
