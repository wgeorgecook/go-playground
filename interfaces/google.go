package main

import (
	"fmt"
)
type gmail struct {
	_type providerType
}

func NewGmail() gmail {
	return gmail{
		_type: GMAIL,
	}
}

func (g gmail) GetType() providerType {
	return GMAIL
}

func (g gmail) SendEmail(message string) error {
	fmt.Printf("%v\n", message)
	return nil
}