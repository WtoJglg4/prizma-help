package server

import (
	"context"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	api "github.com/WtoJglg4/prizma-help/api/service"
)

type Server struct {
}

func (Server) GetGet(ctx context.Context) (api.GetGetRes, error) {
	fmt.Println("Got request")
	rand.NewSource(time.Now().UnixNano())
	servicesNumber := rand.Intn(10) + 1 // Intn(10) дает 0-9, +1 → 1-10
	services := make([]api.Service, 0, servicesNumber)
	for i := range servicesNumber {
		ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255)+1, rand.Intn(255)+1, rand.Intn(255)+1, rand.Intn(255)+1)
		fmt.Printf("Service: %d, ip: %s\n", i+1, ip)
		services = append(services, api.Service{
			CreateIndex: fmt.Sprintf("%d", rand.Intn(10000000)),
			Key:         fmt.Sprintf("service-%d", i+1),
			Value:       base64.RawStdEncoding.EncodeToString([]byte(ip)),
		})
	}
	fmt.Println()
	return &api.GetResponse{
		Services: services,
	}, nil
}
