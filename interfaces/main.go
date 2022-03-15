package main

import (
	"fmt"
)

func buildProvidersList() []sender {
	return []sender{NewOffice().withAddress("real@outlook.com"), NewGmail()}
}

func main() {
	fmt.Println("Hello!")
	defer fmt.Println("Goodbye.")

	senders := buildProvidersList()
	for _, sendr := range senders {
		fmt.Printf("Sending %v email\n", sendr)
		if err := sendr.SendEmail(fmt.Sprintf("Hello from %v!", sendr)); err != nil {
			fmt.Printf("error sending email: %v\n", err)
			break
		}
		fmt.Println("email sent!")
	}
}
