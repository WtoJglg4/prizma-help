// Пакет run отвечает за инициализацию и запуск всех компонентов сервиса
package run

import (
	"context"
	"fmt"
	"log/slog"

	api "github.com/GdeTutMute/summer_practice/servers/api/signals"
	"github.com/GdeTutMute/summer_practice/servers/internal/httpsrv"
	"github.com/GdeTutMute/summer_practice/servers/internal/signalserver/config"
	"github.com/GdeTutMute/summer_practice/servers/internal/signalserver/server"
)

// Options содержит параметры, необходимые для запуска сервиса
type Options struct {
	// Cfg - конфигурация приложения
	Cfg config.Config
	// Logger - настроенный логгер для записи событий
	Logger *slog.Logger
}

// Run инициализирует и запускает все компоненты сервиса
// Функция блокируется до получения сигнала завершения (SIGINT/SIGTERM)
func Run(ctx context.Context, opts Options) error {
	cfg := opts.Cfg
	logger := opts.Logger

	// Создаем HTTP сервер с указанным адресом и настроенным логгером
	httpSrv := httpsrv.New(httpsrv.Options{
		Addr:   cfg.HTTPAddr,
		Logger: logger.With("logging-entity", "httpsrv"),
	})

	// Создаем обработчик API запросов
	handler, err := api.NewServer(&server.Server{})
	if err != nil {
		return fmt.Errorf("create api handler: %w", err)
	}

	// Регистрируем обработчик на путь "/get"
	httpSrv.Register("/get", handler)

	// Запускаем HTTP сервер
	if err := httpSrv.Start(ctx); err != nil {
		return fmt.Errorf("httpsrv start: %w", err)
	}

	// Ожидаем сигнала завершения (SIGINT или SIGTERM)
	<-ctx.Done()

	// Graceful shutdown: корректно останавливаем HTTP сервер
	return httpSrv.Stop(ctx)
}
