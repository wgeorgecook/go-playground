package main

import (
	"fmt"
)

type gmail struct {
	Provider
}

func (g gmail) String() string {
	return "Gmail"
}

func NewGmail() gmail {
	return gmail{}
}

func (g gmail) withAddress(s string) gmail {
	return gmail{Provider{Address: s}}
}

func (g gmail) SendEmail(message string) error {
	if g.Address == "" {
		return noAddressError
	}
	fmt.Printf("%v\n", message)
	return nil
}
