package main

import (
	"fmt"
)
type office struct {
	_type providerType
}

func NewOffice() office {
	return office{
		_type: OFFICE,
	}
}

func (o office) GetType() providerType {
	return o._type
}

func (o office) SendEmail(message string) error {
	fmt.Printf("%v\n", message)
	return nil
}