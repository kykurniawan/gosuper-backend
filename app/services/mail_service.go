package services

import (
	"bytes"
	"gosuper/app/integrations/mail"
	"gosuper/config"
	"html/template"

	"gopkg.in/gomail.v2"
)

type MailService struct {
	dialer *gomail.Dialer
}

func NewMailService() *MailService {
	return &MailService{
		dialer: mail.CreateDialer(),
	}
}

func (service *MailService) SendMail(
	to string,
	subject string, mailTemplate string,
	data interface{},
) error {
	template, err := template.New("mail").Parse(mailTemplate)
	if err != nil {
		return err
	}

	var bodyBuffer bytes.Buffer

	err = template.Execute(&bodyBuffer, data)

	if err != nil {
		return err
	}

	html := bodyBuffer.String()

	mail := gomail.NewMessage()

	mail.SetHeader("From", config.Mail.MailFrom)
	mail.SetHeader("To", to)
	mail.SetHeader("Reply-To", config.Mail.ReplyTo)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", html)

	return service.dialer.DialAndSend(mail)
}
