// Copyright 2018 The QOS Authors

// Package config moon配置
package config

import "path/filepath"

var (
	DefaultQMoonDir  = ".qmoon"
	defaultConfigDir = "config"
	defaultDataDir   = "data"

	defaultConfigFileName = "config.toml"

	defaultConfigFilePath = filepath.Join(defaultConfigDir, defaultConfigFileName)
)

var cnf *Config

// Config moon配置
type Config struct {
	BaseConfig `mapstructure:",squash"`
	HttpServer *HttpServerConfig `mapstructure:"http"`
	DB         *DBConfig         `mapstructure:"db"`
}

// DefaultConfig returns a default configuration
func DefaultConfig() *Config {
	return &Config{
		BaseConfig: DefaultBaseConfig(),
		HttpServer: DefaultRPCConfig(),
		DB:         DefaultDBConfig(),
	}
}

// TestConfig returns a configuration that can be used for testing
func TestConfig() *Config {
	return &Config{
		BaseConfig: TestBaseConfig(),
		HttpServer: TestRPCConfig(),
		DB:         TestDBConfig(),
	}
}
