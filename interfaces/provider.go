package main

type sender interface {
	SendEmail(string) error
}

type Provider struct {
	sender
	Address string
}

type providerError string

func (p providerError) Error() string {
	return string(p)
}

const noAddressError = providerError("no address set")
