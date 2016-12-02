package Figo

import (
	"net/smtp"
	"strings"
)

type SMTPClient struct {
	user     string
	password string
	host     string
}

func GetSMTPClient(user, password, host string) *SMTPClient {
	return &SMTPClient{
		user:     user,
		password: password,
		host:     host,
	}
}

func (p *SMTPClient) Send(to, subject, body string) error {
	auth := smtp.PlainAuth("", p.user, p.password, strings.Split(p.host, ":")[0])
	var content_type = "Content-Type: text/html; charset=UTF-8"
	msg := []byte("To: " + to + "\r\nFrom: " + p.user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	return smtp.SendMail(p.host, auth, p.user, send_to, msg)
}
