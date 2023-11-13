package mail

import (
	"crypto/tls"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func CreateDialer() *gomail.Dialer {
	host := os.Getenv("MAIL_HOST")
	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		panic(err)
	}
	username := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")

	dialer := gomail.NewDialer(host, port, username, password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return dialer
}
