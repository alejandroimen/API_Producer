package adapters

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	amqp "github.com/streadway/amqp"
	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/domain/entities"
	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/domain/repository"
)

var conn *amqp.Connection
var channel *amqp.Channel

func InitRabbitMQ() {
	
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar el archivo .env")
	}

	username := os.Getenv("RABBITMQ_USER")
	password := os.Getenv("RABBITMQ_PASS")
	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")

	rabbitURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, host, port)
	
	var errConn error
	conn, errConn = amqp.Dial(rabbitURL)
	if errConn != nil {
		log.Printf("Error al conectar con RabbitMQ: %s. Reintentando...", errConn)
		time.Sleep(5 * time.Second)
		conn, errConn = amqp.Dial(rabbitURL)
		if errConn != nil {
			log.Fatalf("Error      definitivo al conectar con RabbitMQ: %s", errConn)
		}
	}

	channel, err = conn.Channel()
	if err != nil {
		log.Fatalf("Error al abrir un canal en RabbitMQ: %s", err)
	}

	err = channel.ExchangeDeclare(
		"orders_exchange", // nombre del exchange
		"direct",          // tipo
		true,              // durable
		false,             // auto-deleted
		false,             // internal
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		log.Fatalf("Error al declarar el exchange: %s", err)
	}

	queues := []string{"created.order", "order.confirmed", "order.rejected"}
	for _, queue := range queues {
		_, err = channel.QueueDeclare(
			queue,
			true,  // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
		)
		if err != nil {
			log.Fatalf("Error al declarar la cola '%s': %s", queue, err)
		}

		err = channel.QueueBind(
			queue,           // nombre de la cola
			queue,           // routing key
			"orders_exchange", // exchange
			false,           // no-wait
			nil,             // arguments
		)
		if err != nil {
			log.Fatalf("Error al hacer binding de la cola '%s': %s", queue, err)
		}
	}

	log.Println("Conectado a RabbitMQ exitosamente")
}

func CloseRabbitMQ() {
	if channel != nil {
		channel.Close()
	}
	if conn != nil {
		conn.Close()
	}
}

func GetChannel() *amqp.Channel {
	return channel
}

func ConsumeConfirmedOrders(repo repository.OrderRepository) {
	msgs, err := channel.Consume(
		"order.confirmed",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Error al registrar el consumidor para 'order.confirmed': %s", err)
	}

	log.Println("Consumidor de 'order.confirmed' iniciado correctamente")

	go func() {
		for d := range msgs {
			var order entities.Order
			if err := json.Unmarshal(d.Body, &order); err != nil {
				log.Printf("Error al decodificar el mensaje de 'order.confirmed': %s", err)
				continue
			}

			log.Printf("Mensaje recibido en 'order.confirmed': %+v", order)

			// Actualizar el estado del pedido a "Enviado"
			order.Estado = "Enviado"
			if err := repo.Update(order); err != nil {
				log.Printf("Error al actualizar el pedido %d: %s", order.Id, err)
				continue
			}

			log.Printf("Pedido %d confirmado y actualizado a 'Enviado'", order.Id)
		}
	}()
}

func ConsumeRejectedOrders(repo repository.OrderRepository) {
	msgs, err := channel.Consume(
		"order.rejected",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Error al registrar el consumidor para 'order.rejected': %s", err)
	}

	log.Println("Consumidor de 'order.rejected' iniciado correctamente")

	go func() {
		for d := range msgs {
			var order entities.Order
			if err := json.Unmarshal(d.Body, &order); err != nil {
				log.Printf("Error al decodificar el mensaje de 'order.rejected': %s", err)
				continue
			}

			log.Printf("Mensaje recibido en 'order.rejected': %+v", order)

			// Actualizar el estado del pedido a "Fallido"
			order.Estado = "Fallido"
			if err := repo.Update(order); err != nil {
				log.Printf("Error al actualizar el pedido %d: %s", order.Id, err)
				continue
			}

			log.Printf("Pedido %d rechazado y actualizado a 'Fallido'", order.Id)
		}
	}()
}