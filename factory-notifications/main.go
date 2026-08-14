package main

import (
	"factory-method-notifications/internal/factory"
	"fmt"
	"log"
)

func processEvent(channelType, message, user string) {
	fmt.Printf("\n--- Procesando evento para canal: %s ---\n", channelType)

	notifier, err := factory.GetNotifier(channelType)
	if err != nil {
		log.Printf("Error al generar notificador: %v\n", err)
		return
	}

	err = notifier.Send(message, user)
	if err != nil {
		log.Printf("Error en el envío: %v\n", err)
	}
}

func main() {
	fmt.Println("Iniciando Sistema de Notificaciones Multicanal")

	processEvent(factory.TypeEmail, "Bienvenido a la plataforma, tu registro fue exitoso.", "jaicel@ejemplo.com")
	processEvent(factory.TypeSMS, "Alerta: Nuevo inicio de sesión detectado.", "+5917123543")
	processEvent(factory.TypePush, "Tienes tareas pendientes para hoy.", "user_device_id_8932")
	processEvent("telegram", "Mensaje de prueba", "@usuario")
}
