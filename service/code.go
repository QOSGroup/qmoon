// Copyright 2018 The QOS Authors

package service

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/db/model"
	"github.com/QOSGroup/qmoon/lib/mail"
	"github.com/QOSGroup/qmoon/utils"
)

const VerifyCodeExpired = time.Minute * 15
const VerifyCodeDuration = time.Minute

// CheckCode 验证码校验
func CheckCode(email, code string) bool {
	vc, err := model.VerifyCodeByEmail(db.Db, utils.NullString(email))
	if err != nil {
		return false
	}

	now := time.Now()

	if vc.CreatedAt.Time.Add(VerifyCodeExpired).Before(now) {
		return false
	}

	vc.Delete(db.Db)

	return vc.Code.String == code
}

func mailMsg(from string, to []string, code string) []byte {
	nickname := "admin"
	user := from
	subject := "QOS验证码"
	contentType := "Content-Type: text/plain; charset=UTF-8"
	body := fmt.Sprintf(`
	亲爱的用户：
		您好!

		您本次操作验正码为：%s, %d分钟内有效, 若非本人操作，请忽略!
	`, code, int(VerifyCodeExpired.Minutes()))
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)

	return msg
}

// SendCode 发送短信验证码
func SendCode(mailSmtpServer, mailUser, mailPassword, email string) error {
	now := time.Now()
	code := utils.RandmonCode(6)

	vc, err := model.VerifyCodeByEmail(db.Db, utils.NullString(email))
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
		vc = &model.VerifyCode{}
		vc.Email = utils.NullString(email)
		vc.Code = utils.NullString(code)
		vc.CreatedAt = utils.NullTime(now)

		if err := vc.Insert(db.Db); err != nil {
			return err
		}
	} else {
		if vc.CreatedAt.Time.Add(VerifyCodeDuration).After(now) {
			return errors.New("验证请求太频繁，请稍后再试")
		}
		vc.Email = utils.NullString(email)
		vc.Code = utils.NullString(code)
		vc.CreatedAt = utils.NullTime(now)

		if err := vc.Update(db.Db); err != nil {
			return err
		}
	}

	msg := mailMsg(mailUser, []string{email}, code)
	err = mail.New(mailSmtpServer, mailUser, mailPassword).Send([]string{email}, msg)
	if err != nil {
		return err
	}

	return nil
}
