package queue

type QueueConsumerInterface interface {
	GetRoutingKey() string
	GetQueueName() string
	GetConsumerName() string
	Process(body []byte) (interface{}, error)
	OnSuccess(res interface{})
	OnFailure(err error)
}
