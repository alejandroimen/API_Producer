package services

import (
    "log"

    "github.com/alejandroimen/API_Producer/src/users/domain/services"
    "github.com/alejandroimen/API_Producer/src/users/infrastructure/adapters"
)

func InitRabbitMQ(connectionString string) services.RabbitMQService {
    service, err := adapters.NewRabbitMQAdapter(connectionString)
    if err != nil {
        log.Fatalf("Error inicializando RabbitMQ: %s", err)
    }
    return service
}
