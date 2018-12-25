// Copyright 2018 The QOS Authors

package hadmin

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const sendCodeUrl = "/admin/sendCode"

// SendCodeGinSendCode 注册
func SendCodeGinSendCode(r *gin.Engine) {
	r.POST(sendCodeUrl, sendCodeGin())
}

type sendCodeQuery struct {
	Email string `json:"email"`
}

func (q sendCodeQuery) Validator() error {
	if q.Email == "" {
		return errors.New("email不能为空")
	}

	return nil
}

var emailSmtpServer string
var emailUser string
var emailPassword string

func init() {
	emailSmtpServer = os.Getenv("MailSmtpServer")
	emailUser = os.Getenv("MailUser")
	emailPassword = os.Getenv("MailPassword")
}

func (q *sendCodeQuery) ServerHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("aaaa"))
}

func sendCodeGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqObj sendCodeQuery
		if err := c.ShouldBind(&reqObj); err != nil {
			c.JSON(http.StatusOK, types.RPCParseError("", err))
			return
		}

		fmt.Printf("req:%+v\n", reqObj)
		if err := reqObj.Validator(); err != nil {
			c.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
			return
		}

		logrus.WithField("model", "hadmin").
			WithField("MailSmtpServer", emailSmtpServer).
			WithField("MailUser", emailUser).
			WithField("MailPassword", emailPassword).
			Debug()
		if err := service.SendCode(emailSmtpServer, emailUser, emailPassword, reqObj.Email); err != nil {
			c.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
			return
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", nil))
	}
}
