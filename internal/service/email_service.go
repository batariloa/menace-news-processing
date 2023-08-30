package service

import (
	"net/smtp"

	"github.com/batariloa/bobber/internal/datastruct"
	"gopkg.in/gomail.v2"
)

type EmailService struct {
	SenderEmail    string
	SenderPassword string
	SMTPHost       string
	SMTPPort       int
}

func NewEmailService(senderEmail, senderPassword, smtpHost string, smtpPort int) *EmailService {
	return &EmailService{
		SenderEmail:    senderEmail,
		SenderPassword: senderPassword,
		SMTPHost:       smtpHost,
		SMTPPort:       smtpPort,
	}
}

func (es *EmailService) SendEmail(email *datastruct.Email) error {
	// Create a new email message
	m := gomail.NewMessage()
	m.SetHeader("From", es.SenderEmail)
	m.SetHeader("To", email.Recipient)
	m.SetHeader("Subject", "Test Email")
	m.SetBody("text/html", "<html><body><h4>"+email.Title+"</h4><p>This man truly knows no bounds.</p></body></html>")

	// Create an SMTP dialer with the provided SMTP server details
	d := gomail.NewDialer(es.SMTPHost, es.SMTPPort, es.SenderEmail, es.SenderPassword)

	// Set the correct SMTP authentication options
	d.Auth = smtp.PlainAuth("", es.SenderEmail, es.SenderPassword, es.SMTPHost)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
