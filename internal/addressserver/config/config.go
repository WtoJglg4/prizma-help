// Пакет config отвечает за конфигурацию приложения
package config

import (
	"os"
)

// Config содержит все настройки приложения
type Config struct {
	// HTTPAddr - адрес и порт для HTTP сервера (например, ":5000")
	HTTPAddr             string
	SignalServerAddr     string
	StatisticsServerAddr string
}

// Load загружает конфигурацию из переменных окружения
// Если переменные не установлены, используются значения по умолчанию
func Load() Config {
	// Создаем конфигурацию со значениями по умолчанию
	cfg := Config{
		HTTPAddr:             "127.0.0.1:5000", // По умолчанию сервер слушает на порту 5000
		SignalServerAddr:     "127.0.0.1:5001", // По умолчанию сервер слушает на порту 5001
		StatisticsServerAddr: "127.0.0.1:5002", // По умолчанию сервер слушает на порту 5002
	}

	// Проверяем, установлена ли переменная окружения HTTP_ADDR
	// Если да - используем её значение вместо значения по умолчанию
	httpAddr, ok := os.LookupEnv("HTTP_ADDR")
	if ok {
		cfg.HTTPAddr = httpAddr
	}
	signalServerAddr, ok := os.LookupEnv("SIGNAL_SERVER_ADDR")
	if ok {
		cfg.SignalServerAddr = signalServerAddr
	}
	statisticsServerAddr, ok := os.LookupEnv("STATISTICS_SERVER_ADDR")
	if ok {
		cfg.StatisticsServerAddr = statisticsServerAddr
	}

	return cfg
}
