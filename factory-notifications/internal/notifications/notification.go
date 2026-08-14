package notification

type Notifier interface {
	Send(message string, user string) error
}
