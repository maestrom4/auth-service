package utils

import (
	"auth-service/internal/types"
	"strings"

	"net/smtp"

	log "github.com/sirupsen/logrus"
)

// Note: move interfaces to another folder later.
type EmailSender interface {
	SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error
}

type SmtpClient struct{}

func (s *SmtpClient) SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	return smtp.SendMail(addr, a, from, to, msg)
}

func SendVerificationEmail(eData types.EmailOpt, sender EmailSender) error {
	from := strings.TrimSpace(eData.EmailFrom)
	password := strings.TrimSpace(eData.Password)
	to := eData.Email
	subject := eData.Message
	body := eData.Body

	smtpHost := "smtp.ethereal.email"
	smtpPort := "587"

	header := map[string]string{
		"From":         from,
		"To":           to,
		"Subject":      subject,
		"Content-Type": "text/html; charset=\"utf-8\"",
	}

	message := ""
	for k, v := range header {
		message += k + ": " + v + "\r\n"
	}
	message += "\r\n" + body

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := sender.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return err
	}

	log.Println("Email sent successfully!")
	return nil
}
