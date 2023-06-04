package services

import (
	"bytes"
	"fmt"
	"glucovie/internal/models"
	smtpgomail "glucovie/pkg/smtp"
	"html/template"
	"time"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string, level []models.GlucoseResponse) error {
	type week struct {
		Luni     string
		Marti    string
		Miercuri string
		Joi      string
		Vineri   string
		Sambata  string
		Duminica string
		Date     string
	}

	var w = week{}
	w.Date = time.Now().Format("01.02.2006")

	for _, v := range level {
		switch v.Day {
		case 0:
			w.Luni = fmt.Sprint(v.Level)
		case 1:
			w.Marti = fmt.Sprint(v.Level)
		case 2:
			w.Miercuri = fmt.Sprint(v.Level)
		case 3:
			w.Joi = fmt.Sprint(v.Level)
		case 4:
			w.Vineri = fmt.Sprint(v.Level)
		case 5:
			w.Sambata = fmt.Sprint(v.Level)
		case 6:
			w.Duminica = fmt.Sprint(v.Level)
		}
	}

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
	if err = t.Execute(buffer, w); err != nil {
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
