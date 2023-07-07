package drivers

import (
	"fmt"
)

type JustCall struct {
	apiKey    string
	apiSecret string
}

func (j *JustCall) Send(body string, to string) {
	// Connect via apiKey and apiSecret...
	fmt.Printf(">>> Sending message: '%s', to email: %s\n", body, to)
}

func NewJustCall(apiKey, apiSecret string) *JustCall {
	return &JustCall{
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
}
