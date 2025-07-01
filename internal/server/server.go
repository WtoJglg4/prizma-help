// Пакет server реализует бизнес-логику сервиса
package server

import (
	"context"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	api "github.com/WtoJglg4/prizma-help/api/service"
)

// Server - структура, реализующая методы API сервиса
type Server struct {
}

// GetGet обрабатывает GET-запросы к эндпоинту /get
// Возвращает случайный список сервисов с их IP-адресами
func (Server) GetGet(ctx context.Context) (api.GetGetRes, error) {
	fmt.Println("Got request")

	// Инициализируем генератор случайных чисел текущим временем
	rand.NewSource(time.Now().UnixNano())

	// Генерируем случайное количество сервисов (от 1 до 10)
	servicesNumber := rand.Intn(10) + 1

	// Создаем слайс для хранения информации о сервисах
	services := make([]api.Service, 0, servicesNumber)

	// Генерируем данные для каждого сервиса
	for i := range servicesNumber {
		// Генерируем случайный IP-адрес (каждый октет от 1 до 255)
		ip := fmt.Sprintf("%d.%d.%d.%d",
			rand.Intn(255)+1,
			rand.Intn(255)+1,
			rand.Intn(255)+1,
			rand.Intn(255)+1,
		)
		fmt.Printf("Service: %d, ip: %s\n", i+1, ip)

		// Добавляем информацию о сервисе в список:
		// - CreateIndex: случайный индекс создания
		// - Key: имя сервиса в формате "service-N"
		// - Value: IP-адрес, закодированный в base64
		services = append(services, api.Service{
			CreateIndex: fmt.Sprintf("%d", rand.Intn(10000000)),
			Key:         fmt.Sprintf("service-%d", i+1),
			Value:       base64.RawStdEncoding.EncodeToString([]byte(ip)),
		})
	}
	fmt.Println()

	// Возвращаем ответ с списком сервисов
	return &api.GetResponse{
		Services: services,
	}, nil
}
