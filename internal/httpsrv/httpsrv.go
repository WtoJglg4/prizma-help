// Пакет httpsrv предоставляет обертку над стандартным HTTP сервером Go
// с дополнительными возможностями логирования и управления жизненным циклом
package httpsrv

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

// Server представляет собой HTTP сервер с расширенной функциональностью
type Server struct {
	router *http.ServeMux // маршрутизатор HTTP запросов
	server *http.Server   // стандартный HTTP сервер Go
	logger *slog.Logger   // структурированный логгер
}

// Options содержит параметры для создания нового HTTP сервера
type Options struct {
	Addr   string       // адрес и порт для прослушивания (например, ":5000")
	Logger *slog.Logger // логгер для записи событий сервера
}

// New создает новый экземпляр HTTP сервера с указанными параметрами
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

// Register регистрирует новый обработчик HTTP запросов по указанному пути
// Позволяет также добавить middleware-обработчики, которые будут выполняться
// перед основным обработчиком в порядке их указания
func (s *Server) Register(
	path string,
	handler http.Handler,
	middlewares ...func(handler http.Handler) http.Handler,
) {
	// Применяем все middleware к основному обработчику
	for _, m := range middlewares {
		handler = m(handler)
	}
	s.router.Handle(path, handler)
}

// RegisterFunc регистрирует функцию-обработчик HTTP запросов по указанному пути
// Аналогично Register, но принимает http.HandlerFunc вместо http.Handler
func (s *Server) RegisterFunc(
	path string,
	handler http.HandlerFunc,
	middlewares ...func(handlerFunc http.HandlerFunc) http.HandlerFunc,
) {
	// Применяем все middleware к функции-обработчику
	for _, m := range middlewares {
		handler = m(handler)
	}
	s.router.Handle(path, handler)
}

// Start запускает HTTP сервер в отдельной горутине
// Возвращает ошибку только если не удалось запустить сервер
func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("starting HTTP server")
	go func() {
		// ListenAndServe блокирует выполнение до ошибки или закрытия сервера
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error("http server ListenAndServe failed", "err", err)
		}
	}()
	return nil
}

// Stop gracefully останавливает HTTP сервер
// Ожидает завершения всех текущих запросов или истечения контекста
func (s *Server) Stop(ctx context.Context) error {
	s.logger.Info("stopping HTTP server on")
	err := s.server.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("http server Shutdown failed: %w", err)
	}
	return nil
}
