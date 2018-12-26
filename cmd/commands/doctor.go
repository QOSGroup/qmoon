// Copyright 2018 The QOS Authors

package commands

import (
	"log"

	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/service"
	"github.com/spf13/cobra"
)

// DoctorCmd 健康检查
var DoctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "健康检查",
	RunE:  doctor,
}

func init() {
	registerFlagsDb(DoctorCmd)
}

func doctor(cmd *cobra.Command, args []string) error {
	err := db.InitDb(config.DB, logger)
	if err != nil {
		log.Printf("Check db fail err:%s", err.Error())
	}

	err = service.Doctor()
	if err != nil {
		log.Printf("Check env fail err:%s", err.Error())
	}

	return nil
}
