// Copyright 2018 The QOS Authors

package commands

import (
	"log"
	"os"

	"github.com/QOSGroup/qmoon/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// DoctorCmd 健康检查
var DoctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "健康检查",
	RunE:  doctor,
}

var (
	email           string
	emailSmtpServer string
	emailUser       string
	emailPassword   string
)

func init() {
	emailSmtpServer = os.Getenv("MailSmtpServer")
	emailUser = os.Getenv("MailUser")
	emailPassword = os.Getenv("MailPassword")
	email = "xxxx@163.com"
	registerFlagsDb(DoctorCmd)
}

func doctor(cmd *cobra.Command, args []string) error {

	logrus.WithField("model", "hadmin").
		WithField("MailSmtpServer", emailSmtpServer).
		WithField("MailUser", emailUser).
		WithField("MailPassword", emailPassword).
		Info()

	var err error
	//
	//err = service.Doctor()
	//if err != nil {
	//	log.Printf("Check env fail err:%s", err.Error())
	//}

	err = service.SendCode(emailSmtpServer, emailUser, emailPassword, email)
	if err != nil {
		log.Printf("Check mail fail err:%s", err.Error())
	}

	return nil
}
