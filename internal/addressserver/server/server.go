// Пакет server реализует бизнес-логику сервиса
package server

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	api "github.com/WtoJglg4/prizma-help/api/addressserver"
	"github.com/WtoJglg4/prizma-help/internal/addressserver/config"
)

// Server - структура, реализующая методы API сервиса
type Server struct {
	signalServerAddr     string
	statisticsServerAddr string
}

func New(conf config.Config) *Server {
	return &Server{
		signalServerAddr:     conf.SignalServerAddr,
		statisticsServerAddr: conf.StatisticsServerAddr,
	}
}

// GetGet обрабатывает GET-запросы к эндпоинту /get
// Возвращает список сервисов с адресами сигналов и статистики
func (s *Server) GetGet(ctx context.Context) (api.GetGetRes, error) {
	// Инициализируем генератор случайных чисел текущим временем
	rand.NewSource(time.Now().UnixNano())

	// Возвращаем ответ с списком сервисов
	return &api.GetResponse{
		Services: []api.Service{
			{
				ID:    fmt.Sprintf("%d", rand.Intn(10000000)),
				Key:   "signal",
				Value: s.signalServerAddr,
			},
			{
				ID:    fmt.Sprintf("%d", rand.Intn(10000000)),
				Key:   "statistics",
				Value: s.statisticsServerAddr,
			},
		},
	}, nil
}
