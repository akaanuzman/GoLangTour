package services

import (
	"note-app/src/core/config"

	"gopkg.in/gomail.v2"
)

type IMailService interface {
	SendEmail(to string, subject string, body string) error
}

type MailService struct {
	config *config.Config
}

func NewMailService() IMailService {
	cfg := config.GetConfig()
	return &MailService{
		config: cfg,
	}
}

func (mailService *MailService) SendEmail(to string, subject string, body string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", mailService.config.Mail)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		mailService.config.MailHost,
		mailService.config.MailPort,
		mailService.config.Mail,
		mailService.config.MailPassword,
	)

	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}

	return nil
}
