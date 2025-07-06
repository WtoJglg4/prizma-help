// Пакет server реализует бизнес-логику сервиса
package server

import (
	"context"
	"math"
	"math/rand"
	"sync"
	"time"

	api "github.com/GdeTutMute/summer_practice/servers/api/signals"
)

// Server - структура, реализующая методы API сервиса
type Server struct {
	mu           sync.Mutex
	signalsCount int
}

// GetGet обрабатывает GET-запросы к эндпоинту /get
// Возвращает сигнал
func (s *Server) GetGet(ctx context.Context) (api.GetGetRes, error) {
	s.mu.Lock()
	s.signalsCount++
	defer s.mu.Unlock()
	// Инициализируем генератор случайных чисел текущим временем
	rand.NewSource(time.Now().UnixNano())
	pointsNumber := rand.Intn(11) + 20 // от 20 до 30 точек
	x, y := make([]float64, pointsNumber), make([]float64, pointsNumber)
	step := float64(30) / float64(pointsNumber)
	for i := range x {
		x[i] = float64(i) * step
		y[i] = math.Sin(x[i])
	}
	return &api.SignalResponse{
		ID:   s.signalsCount,
		Name: "sinus",
		X:    x,
		Y:    y,
	}, nil
}
