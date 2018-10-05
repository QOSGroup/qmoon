// Copyright 2018 The QOS Authors

package commands

import (
	"os"

	cfg "github.com/QOSGroup/qmoon/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
	tmflags "github.com/tendermint/tendermint/libs/cli/flags"
	"github.com/tendermint/tendermint/libs/log"
)

var (
	config = cfg.DefaultConfig()
	logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))
)

func init() {
	registerFlagsRootCmd(RootCmd)
}

// ParseConfig retrieves the default environment configuration,
// sets up the Tendermint root and ensures that the root exists
func ParseConfig() (*cfg.Config, error) {
	conf := cfg.DefaultConfig()
	err := viper.Unmarshal(conf)
	if err != nil {
		return nil, err
	}
	return conf, err
}

func registerFlagsRootCmd(cmd *cobra.Command) {
	cmd.PersistentFlags().String("log_level", config.LogLevel, "Log level")
}

func registerFlagsDb(cmd *cobra.Command) {
	cmd.Flags().String("db.driver", config.DB.DriverName, "db driver type: postgres, mysql")
	cmd.Flags().String("db.user", config.DB.User, "db user")
	cmd.Flags().String("db.password", config.DB.Password, "db password")
	cmd.Flags().String("db.addr", config.DB.Addr, "db addr")
	cmd.Flags().String("db.database", config.DB.Database, "db database")
}

func registerFlagsHttpServer(cmd *cobra.Command) {
	cmd.Flags().String("http.laddr", config.HttpServer.ListenAddress, "http server listen addr")
}

// RootCmd moon主命令
var RootCmd = &cobra.Command{
	Use:   "qmoon",
	Short: "qmoon server",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		if cmd.Name() == VersionCmd.Name() {
			return nil
		}
		config, err = ParseConfig()
		if err != nil {
			return err
		}
		logger, err = tmflags.ParseLogLevel(config.LogLevel, logger, cfg.DefaultLogLevel())
		if err != nil {
			return err
		}
		if viper.GetBool(cli.TraceFlag) {
			logger = log.NewTracingLogger(logger)
		}
		logger = logger.With("module", "main")

		return nil
	},
}
