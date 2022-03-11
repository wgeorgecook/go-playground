package main

import (
	"fmt"
)

func buildProvidersList() []provider {
	return []provider{NewOffice(), NewGmail()}
}

func main() {
	fmt.Println("Hello!")
	defer fmt.Println("Goodbye.")

	providers := buildProvidersList()
	for _, provider := range providers {
		fmt.Printf("Sending %v email\n", provider.GetType())
		if err := provider.SendEmail("test email to send"); err != nil {
			fmt.Printf("error sending email: %v", err)
			break
		}
		fmt.Println("email sent!")
	}
}
