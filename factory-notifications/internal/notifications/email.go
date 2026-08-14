package notification

import (
	"fmt"
	"log"
)

type emailNotifier struct {
	smtpServer string
}

func NewEmailNotifier() *emailNotifier {
	return &emailNotifier{
		smtpServer: "smtp.sendgrid.net",
	}
}

func (e *emailNotifier) Send(message string, user string) error {
	log.Printf("[EMAIL] Conectando a %s...", e.smtpServer)
	fmt.Printf("--> [EMAIL] Enviando mensaje: '%s' a la dirección: %s\n", message, user)
	return nil
}
