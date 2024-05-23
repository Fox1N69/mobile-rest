package service

import "gopkg.in/gomail.v2"

type EmailService struct {
	Dialer *gomail.Dialer
}

func NewEmailService() *EmailService {
	dialer := gomail.NewDialer("smtp.gmail.com", 587, "diogoc1707@gmail.com", "tywj kisc qoji bqdc")
	return &EmailService{
		Dialer: dialer,
	}
}

func (s *EmailService) SendMail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "diogoc1707@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	return s.Dialer.DialAndSend(m)
}
