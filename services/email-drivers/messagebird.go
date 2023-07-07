package drivers

import (
	"fmt"
)

type MessageBird struct {
	host string
	port int
}

func (m *MessageBird) Send(body string, to string) {
	// Connect via host and port...
	fmt.Printf(">>> Sending message: '%s', to email: %s\n", body, to)
}

func NewMessageBird(host string, port int) *MessageBird {
	return &MessageBird{
		host: host,
		port: port,
	}
}
