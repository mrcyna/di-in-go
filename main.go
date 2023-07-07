package main

import (
	"di/factories"
	"di/services"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Unable to start the application without env variables")
	}

	emailDriver, err := factories.EmailDriverName(os.Getenv("EMAIL_DRIVER"))
	if err != nil {
		panic(err)
	}

	emailService, _ := factories.EmailService(emailDriver)

	subscribers := []services.Subscriber{
		{
			FirstName: "Sina",
			LastName:  "Ahmadpour",
			Email:     "sina@gmail.test",
		},
		{
			FirstName: "Behzad",
			LastName:  "FazelAsl",
			Email:     "behzad@gmail.test",
		},
		{
			FirstName: "Saeid",
			LastName:  "Taheri",
			Email:     "saeid@gmail.test",
		},
	}
	services.NewNewsletter(&subscribers, &emailService).Announce("New tutorial course released!")
}
