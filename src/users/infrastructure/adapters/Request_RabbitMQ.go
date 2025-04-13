package adapters

import (
    "encoding/json"
    "fmt"
    "log"
    
    "time"
    "github.com/alejandroimen/API_Producer/src/users/domain/services"
    amqp "github.com/streadway/amqp"
)

type RabbitMQAdapter struct {
    conn    *amqp.Connection
    channel *amqp.Channel
}

// Constructor para RabbitMQAdapter
func NewRabbitMQAdapter(connectionString string) (services.RabbitMQService, error) {
    
    conn, err := amqp.Dial(connectionString)
    if err != nil {
        log.Printf("Error al conectar con RabbitMQ: %s. Reintentando...", err)
        time.Sleep(5 * time.Second)
        conn, err = amqp.Dial(connectionString)
        if err != nil {
            return nil, fmt.Errorf("Error definitivo al conectar con RabbitMQ: %w", err)
        }
    }

    channel, err := conn.Channel()
    if err != nil {
        return nil, fmt.Errorf("Error al abrir un canal en RabbitMQ: %w", err)
    }

    err = channel.ExchangeDeclare(
        "citas", // nombre del exchange
        "direct",          // tipo
        true,              // durable
        false,             // auto-deleted
        false,             // internal
        false,             // no-wait
        nil,               // arguments
    )
    if err != nil {
        return nil, fmt.Errorf("Error al declarar el exchange: %w", err)
    }

    queues := []string{"citas.pendientes", "citas.confirmadas"}
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
            return nil, fmt.Errorf("Error al declarar la cola '%s': %w", queue, err)
        }

        err = channel.QueueBind(
            queue,           // nombre de la cola
            queue,           // routing key
            "citas", // exchange
            false,           // no-wait
            nil,             // arguments
        )
        if err != nil {
            return nil, fmt.Errorf("Error al hacer binding de la cola '%s': %w", queue, err)
        }
    }

    log.Println("Conectado a RabbitMQ exitosamente")
    return &RabbitMQAdapter{conn: conn, channel: channel}, nil
}

func (r *RabbitMQAdapter) PublishCreatedUser(idUser int) error {
    if r.conn == nil || r.channel == nil || r.conn.IsClosed() {
        log.Println("Conexión cerrada, intentando reconectar...")
        err := r.reconnect("amqp://rabbit:rabbit@35.170.173.77:5672/vh")
        if err != nil {
            return fmt.Errorf("Error reconectando a RabbitMQ: %w", err)
        }
    }

    idJSON, err := json.Marshal(idUser)
    if err != nil {
        return fmt.Errorf("Error al convertir el ID de usuario a JSON: %w", err)
    }

    err = r.channel.Publish(
        "citas",           // usar el exchange correcto
        "created.user",    // routing key
        false,
        false,
        amqp.Publishing{
            ContentType: "application/json",
            Body:        idJSON,
        },
    )
    if err != nil {
        return fmt.Errorf("Error al publicar el mensaje en RabbitMQ: %w", err)
    }

    log.Printf("📬 IdUser %d publicado en la cola 'created.user'", idUser)
    return nil
}


func (r *RabbitMQAdapter) Close() error {
    if r.channel != nil {
        r.channel.Close()
    }
    if r.conn != nil {
        r.conn.Close()
    }
    log.Println("Conexión con RabbitMQ cerrada")
    return nil
}

func (r *RabbitMQAdapter) reconnect(connectionString string) error {
    conn, err := amqp.Dial(connectionString)
    if err != nil {
        return fmt.Errorf("Error al reconectar con RabbitMQ: %w", err)
    }

    channel, err := conn.Channel()
    if err != nil {
        return fmt.Errorf("Error al reabrir el canal en RabbitMQ: %w", err)
    }

    err = channel.ExchangeDeclare(
        "citas",
        "direct",
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        return fmt.Errorf("Error al declarar el exchange: %w", err)
    }

    queues := []string{"citas.pendientes", "citas.confirmadas"}
    for _, queue := range queues {
        _, err = channel.QueueDeclare(queue, true, false, false, false, nil)
        if err != nil {
            return fmt.Errorf("Error al declarar la cola '%s': %w", queue, err)
        }

        err = channel.QueueBind(queue, queue, "citas", false, nil)
        if err != nil {
            return fmt.Errorf("Error al hacer binding de la cola '%s': %w", queue, err)
        }
    }

    r.conn = conn
    r.channel = channel
    log.Println("🔁 Reconexion exitosa a RabbitMQ")
    return nil
}
func (r *RabbitMQAdapter) StartConsumingCitas() {
    go func() {
        for {
            log.Println("👂 Iniciando consumidor de 'citas.confirmadas' y 'citas.pendientes'...")

            if r.conn == nil || r.conn.IsClosed() {
                log.Println("⚠️ Conexión cerrada. Intentando reconectar...")
                err := r.reconnect("amqp://rabbit:rabbit@35.170.173.77:5672/vh")
                if err != nil {
                    log.Printf("❌ Error al reconectar: %s", err)
                    time.Sleep(5 * time.Second)
                    continue
                }
            }

            // Para monitorear si se cierra algún canal
            closeChan := make(chan *amqp.Error)
            r.conn.NotifyClose(closeChan)

            queues := []string{"citas.confirmadas", "citas.pendientes"}

            for _, queue := range queues {
                // Se crea un canal separado para cada consumidor
                ch, err := r.conn.Channel()
                if err != nil {
                    log.Printf("❌ Error al abrir canal para cola '%s': %s", queue, err)
                    continue
                }

                msgs, err := ch.Consume(
                    queue,
                    "",    // consumer tag
                    true,  // auto-ack
                    false, // exclusive
                    false, // no-local
                    false, // no-wait
                    nil,   // args
                )
                if err != nil {
                    log.Printf("❌ Error al consumir la cola '%s': %s", queue, err)
                    ch.Close() // cerramos canal si no sirve
                    continue
                }

                go func(queue string, ch *amqp.Channel, msgs <-chan amqp.Delivery) {
                    defer ch.Close()

                    for msg := range msgs {
                        var cita map[string]interface{}
                        if err := json.Unmarshal(msg.Body, &cita); err != nil {
                            log.Printf("❌ Error al deserializar mensaje de '%s': %s", queue, err)
                            continue
                        }

                        log.Printf("📥 Cita recibida de la cola '%s': %+v", queue, cita)

                        // Aquí puedes conectar con WebSocket o lógica extra
                    }
                }(queue, ch, msgs)
            }

            // Esperamos a que la conexión se cierre para reconectar todo
            err := <-closeChan
            log.Printf("⚠️ Conexión cerrada con error: %v", err)

            time.Sleep(3 * time.Second)
        }
    }()
}
