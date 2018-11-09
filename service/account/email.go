// Copyright 2018 The QOS Authors

package account

import (
	"net/smtp"
	"os"
	"strings"
)

const (
	from = "xutao@tokenxy.cn"
)

const (
	EmailHostEnv     = "EmailHost"
	EmailUserEnv     = "EmailUser"
	EmailPasswordEnv = "EmailPassword"
)

var (
	emailHost     string
	emailUser     string
	emailPassword string
)

func init() {
	emailHost = os.Getenv(EmailHostEnv) // "smtpdm.aliyun.com"
	//if emailHost == "" {
	//	fmt.Println("EmailHost没有设置环境变量")
	//}
	emailUser = os.Getenv(EmailUserEnv)
	//if emailHost == "" {
	//	fmt.Println("EmailUser没有设置环境变量")
	//}
	emailPassword = os.Getenv(EmailPasswordEnv)
	//if emailHost == "" {
	//	fmt.Println("EmailPassword没有设置环境变量")
	//}
}

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}
func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}
func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		}
	}
	return nil, nil
}

func SendToMail(user, password, host, to, subject, body, mailtype, replyToAddress string) error {
	//hp := strings.Split(host, ":")
	//auth := smtp.PlainAuth("", user, password, hp[0])
	auth := LoginAuth(user, password)
	var contentType string
	if mailtype == "html" {
		contentType = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\nReply-To: " + replyToAddress + "\r\n" + contentType + "\r\n\r\n" + body)

	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendTo, msg)

	return err
}
func send(to string) error {
	subject := "test mail"
	mailtype := "html"
	replyToAddress := ""
	body := `
        <html>
        <body>
        <h3>
        "Test send to email"
        </h3>
        </body>
        </html>
        `
	err := SendToMail(emailUser, emailPassword, emailHost, to, subject, body, mailtype, replyToAddress)
	return err
}

func sendAuthEmail(email ...string) error {
	//m := gomail.NewMessage()
	//m.SetHeader("From", from)
	//m.SetHeader("To", email...)
	////m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	//m.SetHeader("Subject", "欢迎注册QOS")
	//m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	////m.Attach("/home/Alex/lolcat.jpg")
	//
	//d := gomail.NewDialer(emailHost, emailPort, emailUser, emailPassword)
	//
	//if err := d.DialAndSend(m); err != nil {
	//	return err
	//}

	return nil
}

func sendRestPasswordEmail(email string) error {

	return nil
}
