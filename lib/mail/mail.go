// Copyright 2018 The QOS Authors

package mail

import (
	"net/smtp"
)

type client struct {
	addr     string
	username string
	password string
	host     string
}

func New(host, username, password string) *client {
	c := &client{
		host:     host,
		addr:     host + ":25",
		username: username,
		password: password,
	}

	return c
}

func (c *client) Send(to []string, data []byte) error {
	auth := smtp.PlainAuth("", c.username, c.password, c.host)
	err := smtp.SendMail(c.addr, auth, c.username, to, data)

	return err
}
