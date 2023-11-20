package consumers

import (
	"encoding/json"
	"gosuper/app/libs/queue"
	"gosuper/app/services"
	"gosuper/config"
	"log"
)

type SendEmailQueueConsumer struct {
	mailService *services.MailService
}

func NewSendEmailQueueConsumer(mailService *services.MailService) *SendEmailQueueConsumer {
	return &SendEmailQueueConsumer{
		mailService,
	}
}

func (consumer *SendEmailQueueConsumer) GetRoutingKey() string {
	return config.Queue.Mail.RoutingKey
}

func (consumer *SendEmailQueueConsumer) GetQueueName() string {
	return config.Queue.Mail.QueueName
}

func (consumer *SendEmailQueueConsumer) GetConsumerName() string {
	return config.Queue.Mail.ConsumerName
}

func (consumer *SendEmailQueueConsumer) Process(body []byte) (interface{}, error) {
	var queue queue.SendEmailQueue

	err := json.Unmarshal(body, &queue)

	if err != nil {
		return nil, err
	}

	queueData := queue.Data.(map[string]interface{})

	err = consumer.mailService.SendMail(queue.Email, queue.Subject, queue.MailTemplate, queueData)

	if err != nil {
		return nil, err
	}

	return "email sent", nil
}

func (consumer *SendEmailQueueConsumer) OnSuccess(res interface{}) {
	log.Println("success", res)
}

func (consumer *SendEmailQueueConsumer) OnFailure(err error) {
	log.Println("failure", err)
}
