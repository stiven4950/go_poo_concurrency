package main

import "fmt"

// Envío de notificaciones vía Email y SMS

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotification struct{}

func (SMSNotification) SendNotification() {
	fmt.Println("Send Notification vía SMS")
}
func (SMSNotification) GetSender() ISender {
	return SMSNotificatinSender{}
}

type SMSNotificatinSender struct{}

func (SMSNotificatinSender) GetSenderMethod() string {
	return "SMS"
}
func (SMSNotificatinSender) GetSenderChannel() string {
	return "Twilio"
}

type EmailNotification struct{}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending notifictions vía Email")
}
func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct{}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}
func (EmailNotificationSender) GetSenderChannel() string {
	return "Brevo antes SendinBlue"
}

func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}
	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("No notification Type")
}

func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main3() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	sendNotification(smsFactory)
	sendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}
