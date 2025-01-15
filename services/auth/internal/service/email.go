package service

import (
	"github.com/Ruletk/GoMarketplace/pkg/logging"
	"gopkg.in/gomail.v2"
)

type EmailService interface {
	sendEmail(to string, subject string, body string) error
	SendVerificationEmail(email string, token string) error
	SendPasswordResetEmail(email string, token string) error
}

type emailService struct {
	dialer *gomail.Dialer
}

func NewEmailService(dialer *gomail.Dialer) EmailService {
	return &emailService{
		dialer: dialer,
	}
}

func (e emailService) sendEmail(to string, subject string, body string) error {
	logging.Logger.Info("Sending email to: ", to, " with subject: ", subject)
	m := gomail.NewMessage()
	m.SetHeader("From", "test@freesmtpservers.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	if err := e.dialer.DialAndSend(m); err != nil {
		logging.Logger.WithError(err).Error("Error sending email to: ", to)
		return err
	}
	logging.Logger.Info("Email sent to: ", to)
	return nil
}

func (e emailService) SendVerificationEmail(email string, token string) error {
	subject := "Verify your email"
	body := "Click the following link to verify your email: <a href='http://localhost/api/v1/auth/verify/" + token + "'>Link</a>"
	return e.sendEmail(email, subject, body)
}

func (e emailService) SendPasswordResetEmail(email string, token string) error {
	subject := "Password reset"
	body := "Click the following link to reset your password: <a href='http://localhost/api/v1/auth/reset-password/" + token + "'>Link</a>"
	return e.sendEmail(email, subject, body)
}
