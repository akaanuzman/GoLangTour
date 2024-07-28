package helpers

import (
	"chat-app/src/core/config"

	"gopkg.in/gomail.v2"
)

// IMailHelper defines the interface for sending emails.
type IMailHelper interface {
	// SendEmail sends an email to the specified recipient with the given subject and body.
	// Returns an error if the email could not be sent.
	SendEmail(to string, subject string, body string) error
}

// MailHelper is a struct that implements the IMailHelper interface.
type MailHelper struct {
	config *config.Config
}

// NewMailHelper creates a new instance of MailHelper with the provided configuration.
// Returns an instance of IMailHelper.
func NewMailHelper(cfg *config.Config) IMailHelper {
	return &MailHelper{
		config: cfg,
	}
}

// SendEmail sends an email using the configuration provided in the MailHelper struct.
// Parameters:
// - to: the recipient's email address.
// - subject: the subject of the email.
// - body: the body of the email in HTML format.
// Returns an error if the email could not be sent.
func (MailHelper *MailHelper) SendEmail(to string, subject string, body string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", MailHelper.config.Mail)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		MailHelper.config.MailHost,
		MailHelper.config.MailPort,
		MailHelper.config.Mail,
		MailHelper.config.MailPassword,
	)

	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}

	return nil
}
