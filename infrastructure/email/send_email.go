package email

import (
	"fmt"
	"net/smtp"
)

func (e *EmailService) SendEmailVerification(email, token string) error {
	auth := smtp.PlainAuth("", e.SMTPUser, e.SMTPPassword, e.SMTPServer)
	to := []string{email}
	subject := "Subject: Email Verification\n"
	body := fmt.Sprintf("Click the link to verify your email: http://localhost:8080/users/verify-email?token=%s&email=%s", token, email)
	msg := []byte(subject + "\n" + body)

	address := fmt.Sprintf("%s:%s", e.SMTPServer, e.SMTPPort)
	return smtp.SendMail(address, auth, e.FromAddress, to, msg)
}
