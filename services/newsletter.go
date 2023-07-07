package services

import (
	"di/contracts"
	"fmt"
)

type Subscriber struct {
	FirstName string
	LastName  string
	Email     string
}

type Newsletter struct {
	subscribers []Subscriber
	emailDriver contracts.EmailInterface
}

func (n *Newsletter) Announce(message string) {
	for _, subscriber := range n.subscribers {
		body := fmt.Sprintf("Dear %s!, %s", subscriber.FirstName, message)
		n.emailDriver.Send(body, subscriber.Email)
	}
}

func NewNewsletter(subscribers *[]Subscriber, emailDriver *contracts.EmailInterface) *Newsletter {
	return &Newsletter{
		subscribers: *subscribers,
		emailDriver: *emailDriver,
	}
}
