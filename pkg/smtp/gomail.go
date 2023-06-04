package smtpgomail

import (
	"crypto/tls"
	"glucovie/pkg/dotenv"
	"strconv"

	"gopkg.in/gomail.v2"
)

func NewEmailConnection() *gomail.Dialer {
	emailCredentials := dotenv.GetEnvironmentVariable("GOMAIL_PORT")
	mailPort, _ := strconv.Atoi(emailCredentials)
	d := gomail.NewDialer(
		dotenv.GetEnvironmentVariable("GOMAIL_HOST"),
		mailPort,
		dotenv.GetEnvironmentVariable("GOMAIL_MAIL"),
		dotenv.GetEnvironmentVariable("GOMAIL_PASSWORD"),
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return d
}
