package main

import "fmt"

type Notifier interface {
	Send(message string) error
}

type NotificationFactory interface {
	CreateNotifier() Notifier
}

type EmailNotifier struct {
	recipient string
}

func (e *EmailNotifier) Send(message string) error {
	fmt.Printf("Sending email with message: %s, %s\n", e.recipient, message)
	return nil
}

type SMSNotifier struct {
	phoneNumber string
}

func (s *SMSNotifier) Send(message string) error {
	fmt.Printf("Sending SMS with message: %s, %s\n", s.phoneNumber, message)
	return nil
}

type PushNotifier struct {
	deviceID string
}

func (p *PushNotifier) Send(message string) error {
	fmt.Printf("Sending push notification with message: %s, %s\n", p.deviceID, message)
	return nil
}

// ------

type EmailFactory struct {
	recipient string
}

func (f *EmailFactory) CreateNotifier() Notifier {
	return &EmailNotifier{recipient: f.recipient}
}

type SMSFactory struct {
	phoneNumber string
}

func (f *SMSFactory) CreateNotifier() Notifier {
	return &SMSNotifier{phoneNumber: f.phoneNumber}
}

type PushFactory struct {
	deviceID string
}

func (f *PushFactory) CreateNotifier() Notifier {
	return &PushNotifier{deviceID: f.deviceID}
}

// ------

func GetNotificationFactory(notificationType, destination string) (NotificationFactory, error) {
	switch notificationType {
	case "email":
		return &EmailFactory{recipient: destination}, nil
	case "sms":
		return &SMSFactory{phoneNumber: destination}, nil
	case "push":
		return &PushFactory{deviceID: destination}, nil
	default:
		return nil, fmt.Errorf("unknown notification type: %s", notificationType)
	}
}

// -----

type NotificationService struct {
	factory NotificationFactory
}

func NewNotificationService(factory NotificationFactory) *NotificationService {
	return &NotificationService{factory: factory}
}

func (s *NotificationService) Notify(message string) error {
	notifier := s.factory.CreateNotifier()
	return notifier.Send(message)
}

// -----

func main() {
	emailFactory := &EmailFactory{recipient: "usuario@example.com"}
	emailService := NewNotificationService(emailFactory)
	if err := emailService.Notify("Hello via Email!"); err != nil {
		fmt.Println("Error:", err)
	}

	smsFactory := &SMSFactory{phoneNumber: "+1234567890"}
	smsFactoryService := NewNotificationService(smsFactory)
	if err := smsFactoryService.Notify("Hello via SMS!"); err != nil {
		fmt.Println("Error:", err)
	}

	pushFactory := &PushFactory{deviceID: "device123"}
	pushService := NewNotificationService(pushFactory)
	if err := pushService.Notify("Hello via Push Notification!"); err != nil {
		fmt.Println("Error:", err)
	}

	preferences := []struct {
		notificationType string
		destination      string
		message          string
	}{
		{"email", "admin@example.com", "System Alert via Email!"},
		{"sms", "+1987654321", "System Alert via SMS!"},
		{"push", "device456", "System Alert via Push Notification!"},
	}

	for _, pref := range preferences {
		factory, _ := GetNotificationFactory(pref.notificationType, pref.destination)
		if factory != nil {
			service := NewNotificationService(factory)
			if err := service.Notify(pref.message); err != nil {
				fmt.Println("Error:", err)
			}
		}
	}
}
