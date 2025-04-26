package broker

import (
	"log"

	"github.com/veyselaksin/gomailer/pkg/mailer"
)

type EmailTask struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type EmailBroker struct {
	tasks    chan EmailTask
	mailAuth *mailer.Authentication
}

func NewEmailBroker(mailAuth *mailer.Authentication) *EmailBroker {
	broker := &EmailBroker{
		tasks:    make(chan EmailTask, 100),
		mailAuth: mailAuth,
	}
	go broker.run()
	return broker
}

func (b *EmailBroker) run() {
	for task := range b.tasks {
		b.handleTask(task)
	}
}

func (b *EmailBroker) handleTask(task EmailTask) {
	sender := mailer.NewPlainAuth(b.mailAuth)
	msg := mailer.NewMessage(task.Subject, task.Body)
	msg.SetTo([]string{task.To})

	if err := sender.SendMail(msg); err != nil {
		log.Printf("failed to send email to %s: %v", task.To, err)
		return
	}
	log.Printf("email sent successfully to: %s", task.To)
}

func (b *EmailBroker) Publish(task EmailTask) {
	select {
	case b.tasks <- task:
	default:
		log.Printf("email broker queue is full, dropping email to: %s", task.To)
	}
}
