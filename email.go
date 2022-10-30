package Hissec

import "net/smtp"

type Email struct {
	Host     string
	Port     int
	UserName string
	Password string
}

func (email *Email) auth() {
	smtp.PlainAuth("", email.UserName, email.Password, email.Host)
}
