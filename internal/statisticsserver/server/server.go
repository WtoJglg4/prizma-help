// Пакет server реализует бизнес-логику сервиса
package server

import (
	"context"
	"errors"
	"fmt"

	api "github.com/GdeTutMute/summer_practice/servers/api/statistics"
)

var (
	errInvalidRequest = errors.New("invalid request")
)

// Server - структура, реализующая методы API сервиса
type Server struct {
}

func (s *Server) StatisticsPost(ctx context.Context, req *api.StatisticsRequest) (api.StatisticsPostRes, error) {
	fmt.Println("StatisticsPost", req)
	if len(req.Values) == 0 {
		return nil, fmt.Errorf("values array must have at least one element: %w", errInvalidRequest)
	}
	return &api.StatisticsResponse{
		Min:     min(req.Values),
		Max:     max(req.Values),
		Average: average(req.Values),
	}, nil
}

func min(v []float64) float64 {
	m := v[0]
	for _, v := range v {
		if v < m {
			m = v
		}
	}
	return m
}

func max(v []float64) float64 {
	m := v[0]
	for _, v := range v {
		if v > m {
			m = v
		}
	}
	return m
}

func average(v []float64) float64 {
	if len(v) == 1 {
		return v[0]
	}
	sum := 0.0
	for _, v := range v {
		sum += v
	}
	return sum / float64(len(v))
}
