package services

import (
	"bytes"
	smtpgomail "glucovie/pkg/smtp"
	"html/template"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string) error {
	mailDialer := smtpgomail.NewEmailConnection()

	m := gomail.NewMessage()

	m.SetHeader("From", "glucoviesupp@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "GlucoVie Support")

	tempPath := "./templates/email_template.html"
	t, err := template.ParseFiles(tempPath)
	if err != nil {
		return err
	}

	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, nil); err != nil {
		return err
	}

	str := buffer.String()

	m.SetBody("text/html", str)

	err = mailDialer.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil
}
