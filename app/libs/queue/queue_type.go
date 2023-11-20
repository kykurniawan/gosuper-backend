package queue

type SendEmailQueue struct {
	Email        string
	Subject      string
	MailTemplate string
	Data         interface{}
}
