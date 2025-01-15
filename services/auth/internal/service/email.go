package service

import "github.com/Ruletk/GoMarketplace/pkg/logging"

type EmailService interface {
	sendEmail(to string, subject string, body string) error
	SendVerificationEmail(email string, token string) error
	SendPasswordResetEmail(email string, token string) error
}

type emailService struct {
}

func NewEmailService() EmailService {
	return &emailService{}
}

func (e emailService) sendEmail(to string, subject string, body string) error {
	logging.Logger.Info("Sending email to: ", to, " with subject: ", subject)
	return nil
}

func (e emailService) SendVerificationEmail(email string, token string) error {
	subject := "Verify your email"
	body := "Click the following link to verify your email: http://localhost:8080/auth/verify/" + token
	return e.sendEmail(email, subject, body)
}

func (e emailService) SendPasswordResetEmail(email string, token string) error {
	subject := "Password reset"
	body := "Click the following link to reset your password: http://localhost:8080/auth/reset-password/" + token
	return e.sendEmail(email, subject, body)
}
