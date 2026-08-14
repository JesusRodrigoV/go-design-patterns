package notification

import (
	"fmt"
	"log"
)

type pushNotifier struct {
	fcmToken string
}

func NewPushNotifier() *pushNotifier {
	return &pushNotifier{
		fcmToken: "fcm_project_token_xxx",
	}
}

func (p *pushNotifier) Send(message string, user string) error {
	log.Println("[PUSH] Estableciendo conexión con servidor FCM...")
	fmt.Printf("--> [PUSH] Disparando notificación: '%s' al dispositivo del usuario: %s\n", message, user)
	return nil
}
