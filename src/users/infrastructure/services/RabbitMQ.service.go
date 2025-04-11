package services

import (

	adapters "github.com/alejandroimen/API_Producer/src/users/infrastructure/adapters"
	"github.com/alejandroimen/API_Producer/src/users/domain/services"
)

// Inicializar el servicio de BCrypt
func InitRabbitMQ() services. {
	return adapters.NewCorreo()
}