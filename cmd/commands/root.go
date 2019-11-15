// Copyright 2018 The QOS Authors

package commands

import (
	"log"
	"os"
	"path/filepath"

	cfg "github.com/QOSGroup/qmoon/config"
	"github.com/QOSGroup/qmoon/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	HomeFlag = "home"
	config   = cfg.DefaultConfig()
	logger   = log.New(os.Stderr, "", log.LstdFlags)
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
	cmd.PersistentFlags().StringP(HomeFlag, "", config.RootDir, "directory for config and data")
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
	Short: "qmoon cli",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		if cmd.Name() == VersionCmd.Name() {
			return nil
		}

		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}

		homeDir := viper.GetString(HomeFlag)
		viper.Set(HomeFlag, homeDir)
		viper.SetConfigName("config")                         // name of config file (without extension)
		viper.AddConfigPath(homeDir)                          // search root directory
		viper.AddConfigPath(filepath.Join(homeDir, "config")) // search root directory /config

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			// stderr, so if we redirect output to json file, this doesn't appear
			// fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		} else if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// ignore not found error, return other errors
			return err
		}

		config, err = ParseConfig()
		if err != nil {
			return err
		}

		if config.LogLevel == "debug" {
			logrus.SetLevel(logrus.DebugLevel)
		}

		if err := models.InitDb(config.DB); err != nil {
			return err
		}

		//logger.Info("rootCmd", "config.BaseConfig", config.BaseConfig,
		//	"config.HttpServer", config.HttpServer, "config.DB", config.DB)

		return nil
	},
}
