// Copyright 2018 The QOS Authors

package commands

import (
	"fmt"

	"github.com/QOSGroup/qmoon/version"
	"github.com/spf13/cobra"
)

// VersionCmd qmoon 版本
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Version)
	},
}
