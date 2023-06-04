package services

import (
	smtpgomail "glucovie/pkg/smtp"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string) error {
	mailDialer := smtpgomail.NewEmailConnection()

	m := gomail.NewMessage()

	m.SetHeader("From", "glucoviesupp@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "GlucoVie Support")

	m.SetBody("text/html", "Hello World")

	err := mailDialer.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil
}
