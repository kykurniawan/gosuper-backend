package services

import (
	"bytes"
	"crypto/tls"
	"html/template"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type MailService struct {
	dialer *gomail.Dialer
}

func NewMailService() *MailService {
	host := os.Getenv("MAIL_HOST")
	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		panic(err)
	}
	username := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")

	dialer := gomail.NewDialer(host, port, username, password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return &MailService{
		dialer: dialer,
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

	mail.SetHeader("From", os.Getenv("MAIL_FROM"))
	mail.SetHeader("To", to)
	mail.SetHeader("Reply-To", os.Getenv("MAIL_REPLY_TO"))
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", html)

	return service.dialer.DialAndSend(mail)
}
