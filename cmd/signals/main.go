// Основной пакет приложения
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/GdeTutMute/summer_practice/servers/internal/signalserver/config"
	"github.com/GdeTutMute/summer_practice/servers/internal/signalserver/run"
)

// main - точка входа в приложение
func main() {
	// Создаем контекст, который будет отменен при получении сигнала SIGINT (Ctrl+C)
	// или SIGTERM (сигнал завершения процесса)
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	// Запускаем основной процесс приложения с помощью функции run.Run
	// Передаем:
	// - контекст для graceful shutdown
	// - конфигурацию, загруженную из переменных окружения
	// - настроенный JSON-логгер с информацией об источнике логов
	if err := run.Run(ctx, run.Options{
		Cfg: config.Load(),
		Logger: slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{AddSource: true},
			),
		),
	}); err != nil {
		// В случае ошибки выводим её в stderr и завершаем процесс с кодом 1
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
