package mail

import (
	"crypto/tls"
	"gosuper/config"

	"gopkg.in/gomail.v2"
)

func CreateDialer() *gomail.Dialer {
	dialer := gomail.NewDialer(
		config.Mail.Host,
		config.Mail.Port,
		config.Mail.Username,
		config.Mail.Password,
	)

	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return dialer
}
