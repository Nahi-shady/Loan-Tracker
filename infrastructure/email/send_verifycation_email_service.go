package email

import (
	"loan-tracker/domain"
)

type EmailService struct {
	SMTPServer   string
	SMTPPort     string
	SMTPUser     string
	SMTPPassword string
	FromAddress  string
}

func NewEmailService(server, port, user, password, fromAddress string) domain.EmailService {
	return &EmailService{
		SMTPServer:   server,
		SMTPPort:     port,
		SMTPUser:     user,
		SMTPPassword: password,
		FromAddress:  fromAddress,
	}
}
