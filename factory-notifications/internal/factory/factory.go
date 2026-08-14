package factory

import (
	"errors"
	notification "factory-method-notifications/internal/notifications"
	"strings"
)

const (
	TypeEmail = "email"
	TypeSMS   = "sms"
	TypePush  = "push"
)

func GetNotifier(channelType string) (notification.Notifier, error) {
	switch strings.ToLower(channelType) {
	case TypeEmail:
		return notification.NewEmailNotifier(), nil
	case TypeSMS:
		return notification.NewSMSNotifier(), nil
	case TypePush:
		return notification.NewPushNotifier(), nil
	default:
		return nil, errors.New("canal de notificación no soportado: " + channelType)
	}
}
