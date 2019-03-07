// Copyright 2018 The QOS Authors

package service

import (
	"fmt"
	"log"
	"time"

	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/utils"
	"github.com/go-gomail/gomail"
)

// CheckCode 验证码校验
func CheckCode(email, code string) bool {
	vc, err := models.VerifyCodeByEmail(email)
	log.Printf("CheckCode email:%s, code:%s, vc:%+v, err:%v", email, code, vc, err)
	if err != nil {
		return false
	}

	now := time.Now()

	if vc.CreatedAt.Add(models.VerifyCodeExpired).Before(now) {
		return false
	}

	return vc.Code == code
}

// SendCode 发送短信验证码
func SendCode(mailSmtpServer, mailUser, mailPassword, email string) error {
	code := utils.RandmonCode(6)
	_, err := models.CreateVerifyCode(email, code)
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetAddressHeader("From", mailUser, "QOSChain")
	m.SetAddressHeader("To", email, "")
	m.SetHeader("Subject", "验证码")

	contentType := "Content-Type: text/plain; charset=UTF-8"
	body := fmt.Sprintf(`
	亲爱的用户：
		您好!

		您本次操作验正码为：%s, %d分钟内有效, 若非本人操作，请忽略!
	`, code, int(models.VerifyCodeExpired.Minutes()))
	m.SetBody(contentType, body)

	d := gomail.NewDialer(mailSmtpServer, 465, mailUser, mailPassword)
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
