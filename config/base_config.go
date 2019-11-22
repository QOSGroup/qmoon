// Copyright 2018 The QOS Authors

package config

import (
	"os"
	"path/filepath"
)

// BaseConfig defines the base configuration for a Tendermint node
type BaseConfig struct {

	// RootDir 数据根目录
	RootDir string `mapstructure:"home"`

	// LogLevel 日志输出级别 debug info warning error fatal panic
	LogLevel string `mapstructure:"log_level"`
}

// DefaultBaseConfig returns a default base configuration
func DefaultBaseConfig() BaseConfig {
	return BaseConfig{
		RootDir:  os.ExpandEnv(filepath.Join("$HOME", DefaultQMoonDir)),
		LogLevel: DefaultLogLevel(),
	}
}

// TestBaseConfig returns a base configuration
func TestBaseConfig() BaseConfig {
	cfg := DefaultBaseConfig()
	return cfg
}

// DefaultLogLevel returns a default log level of "error"
func DefaultLogLevel() string {
	return "debug"
}
