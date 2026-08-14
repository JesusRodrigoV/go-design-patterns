package notification

import (
	"fmt"
	"log"
)

type smsNotifier struct {
	comunicAPIKey string
}

func NewSMSNotifier() *smsNotifier {
	return &smsNotifier{
		comunicAPIKey: "asdf123456789crlqwe111",
	}
}

func (s *smsNotifier) Send(message string, user string) error {
	log.Println("[SMS] Autenticando con proveedor de telefonía...")
	fmt.Printf("--> [SMS] Enviando SMS: '%s' al número: %s\n", message, user)
	return nil
}
