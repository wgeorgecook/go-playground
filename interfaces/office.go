package main

import (
	"fmt"
)

type office struct {
	Provider
}

func (o office) String() string {
	return "Office"
}

func NewOffice() office {
	return office{}
}

func (o office) withAddress(s string) office {
	return office{Provider{Address: s}}
}

func (o office) SendEmail(message string) error {
	if o.Address == "" {
		return noAddressError
	}
	fmt.Printf("%v\n", message)
	return nil
}
