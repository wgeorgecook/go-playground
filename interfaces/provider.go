package main

type providerType int

const (
	UNDEFINED providerType = iota
	OFFICE
	GMAIL
)

func (p providerType) String() string {
	switch p {
	case OFFICE:
		return "office"
	case GMAIL:
		return "gmail"
	default:
		return "undefined"
	}
}

type provider interface {
	GetType() providerType
	SendEmail(string) error
}
