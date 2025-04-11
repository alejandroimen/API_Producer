package services

type RabbitMQService interface {
    SendCorreo(asunto string, body string)
}