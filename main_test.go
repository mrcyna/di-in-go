package main

import (
	"di/factories"
	"di/services"
	"io"
	"os"
	"strings"
	"testing"
)

func TestAppIsWorking(t *testing.T) {
	// Faking a standard output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Set the env variables
	os.Setenv("MESSAGEBIRD_HOST", "messagebird.local")
	os.Setenv("MESSAGEBIRD_PORT", "1000")

	// Create subscriber
	subscribers := []services.Subscriber{
		{
			FirstName: "john",
			LastName:  "doe",
			Email:     "jobh@app.local",
		},
	}

	// Create EmailDriver
	emailDriverName, _ := factories.EmailDriverName("messagebird")
	emailDriver, _ := factories.EmailService(emailDriverName)

	// Create a Newsletter service
	services.NewNewsletter(&subscribers, &emailDriver).Announce("Test Message")

	// Restoring the original standard output
	w.Close()
	os.Stdout = old
	got, _ := io.ReadAll(r)
	output := string(got)

	if !strings.Contains(output, ">>> Sending message: 'Dear john!, Test Message', to email: jobh@app.local") {
		t.Error("Failure!")
	}
}
