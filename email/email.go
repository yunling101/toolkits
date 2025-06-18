package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type mail struct {
	host     string
	username string
	password string
	port     int
	tls      bool
	message  *gomail.Message
}

func New(host, username, password string, port int) *mail {
	return &mail{
		host:     host,
		username: username,
		password: password,
		port:     port,
		message:  gomail.NewMessage(),
	}
}

func (c *mail) Tls(value bool) *mail {
	c.tls = value
	return c
}

func (c *mail) To(value ...string) *mail {
	c.message.SetHeader("To", value...)
	return c
}

func (c *mail) Subject(subject string) *mail {
	c.message.SetHeader("Subject", subject)
	return c
}

func (c *mail) Body(contentType, body string) *mail {
	c.message.SetBody(contentType, body)
	return c
}

func (c *mail) Attach(filename string) *mail {
	c.message.Attach(filename)
	return c
}

func (c *mail) Deliver() (err error) {
	c.message.SetHeader("From", c.username)
	m := gomail.NewDialer(c.host, c.port, c.username, c.password)
	if c.tls {
		m.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}
	err = m.DialAndSend(c.message)
	return
}
