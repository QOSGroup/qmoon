// Copyright 2018 The QOS Authors

package config

import "os"

/*

 */

// EmailConfig 邮件配置
type EmailConfig struct {
	SmtpServer string
	User       string
	Password   string
}

// DefaultEmailConfig returns a default configuration for the HttpServer server
func DefaultEmailConfig() *EmailConfig {
	return &EmailConfig{
		SmtpServer: os.Getenv("MailSmtpServer"),
		User:       os.Getenv("MailUser"),
		Password:   os.Getenv("MailPassword"),
	}
}

// TestEmailConfig returns a configuration for testing the HttpServer server
func TestEmailConfig() *EmailConfig {
	return &EmailConfig{
		SmtpServer: os.Getenv("MailSmtpServer"),
		User:       os.Getenv("MailUser"),
		Password:   os.Getenv("MailPassword"),
	}
}
