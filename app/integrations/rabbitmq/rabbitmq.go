package rabbitmq

import (
	"gosuper/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

func CreateConnection() (*amqp.Connection, error) {
	return amqp.Dial(config.RabbitMQ.Dsn)
}
