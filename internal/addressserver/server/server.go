// Пакет server реализует бизнес-логику сервиса
package server

import (
	"context"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	api "github.com/GdeTutMute/summer_practice/servers/api/address"
	"github.com/GdeTutMute/summer_practice/servers/internal/addressserver/config"
)

// Server - структура, реализующая методы API сервиса
type Server struct {
	signalServerAddrBase64     string
	statisticsServerAddrBase64 string
}

func New(conf config.Config) *Server {
	return &Server{
		signalServerAddrBase64:     base64.StdEncoding.EncodeToString([]byte(conf.SignalServerAddr)),
		statisticsServerAddrBase64: base64.StdEncoding.EncodeToString([]byte(conf.StatisticsServerAddr)),
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
				Value: s.signalServerAddrBase64,
			},
			{
				ID:    fmt.Sprintf("%d", rand.Intn(10000000)),
				Key:   "statistics",
				Value: s.statisticsServerAddrBase64,
			},
		},
	}, nil
}
