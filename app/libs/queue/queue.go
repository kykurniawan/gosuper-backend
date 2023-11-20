package queue

import (
	"context"
	"encoding/json"
	"gosuper/config"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Queue struct {
	amqp      *amqp.Connection
	consumers []QueueConsumerInterface
}

func NewQueue(amqp *amqp.Connection) *Queue {
	return &Queue{
		amqp,
		[]QueueConsumerInterface{},
	}
}

func (q *Queue) RegisterConsumer(consumer QueueConsumerInterface) {
	q.consumers = append(q.consumers, consumer)
}

func (q *Queue) Publish(routingKey string, queue interface{}) error {
	channel, err := q.amqp.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	ctx := context.Background()

	jsonData, err := json.Marshal(queue)

	if err != nil {
		return err
	}

	body := []byte(jsonData)

	err = channel.PublishWithContext(
		ctx,
		config.Queue.Exchange, // Exchange
		routingKey,            // Routing key
		false,                 // Mandatory
		false,                 // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (q *Queue) Run(ctx context.Context) {
	for _, consumer := range q.consumers {
		go func(consumer QueueConsumerInterface) {
			err := q.consume(
				ctx,
				consumer.GetQueueName(),
				consumer.GetRoutingKey(),
				consumer.GetConsumerName(),
				consumer.Process,
				consumer.OnSuccess,
				consumer.OnFailure,
			)

			if err != nil {
				log.Fatal(err)
			}
		}(consumer)
	}
}

func (q *Queue) consume(
	ctx context.Context,
	queueName,
	routingKey,
	consumerName string,
	process func([]byte) (interface{}, error),
	onSuccess func(interface{}),
	onFailure func(error),
) error {
	channel, err := q.amqp.Channel()

	if err != nil {
		return err
	}

	defer channel.Close()

	queue, err := channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	err = channel.QueueBind(
		queue.Name,
		routingKey,
		config.Queue.Exchange,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	msgs, err := channel.Consume(
		queue.Name,
		consumerName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case d, ok := <-msgs:
			if !ok {
				return nil
			}
			res, err := process(d.Body)

			if err != nil {
				d.Nack(false, true)
				onFailure(err)
			} else {
				d.Ack(false)
				onSuccess(res)
			}
		}
	}
}
