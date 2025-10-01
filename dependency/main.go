package main

import "fmt"

type Notifier interface {
	Send(message string) error
}

type Email struct {
	recipient string
}

func (e *Email) Send(message string) error {
	fmt.Printf("Sending email with message: %s, %s\n", e.recipient, message)
	return nil
}

type SMS struct {
	phoneNumber string
}

func (s *SMS) Send(message string) error {
	fmt.Printf("Sending SMS with message: %s, %s\n", s.phoneNumber, message)
	return nil
}

type NotificationService struct {
	notifier Notifier
}

/*
	NewNotificationService creates a new NotificationService with the given Notifier.

Here we are using the Dependency Injection pattern to pass the Notifier to the NotificationService.
*/
func NewNotificationService(notifier Notifier) *NotificationService {
	return &NotificationService{notifier: notifier}
}

func (s *NotificationService) Notify(message string) error {
	return s.notifier.Send(message)
}

func main() {
	email := Email{recipient: "acetina@example.com"}
	service := NewNotificationService(&email)
	_ = service.Notify("Hello!")

}
