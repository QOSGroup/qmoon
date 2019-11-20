// Copyright 2018 The QOS Authors

package config

import (
	"fmt"
)

// DBConfig
type DBConfig struct {
	DriverName string `mapstructure:"driver"`
	User       string `mapstructure:"user"`
	Password   string `mapstructure:"password"`
	Addr       string `mapstructure:"addr"`
	Database   string `mapstructure:"database"`
}

func (dc *DBConfig) DataSource() string {
	switch dc.DriverName {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s)/%s", dc.User, dc.Password, dc.Addr, dc.Database)
	case "postgres":
		// postgres://postgres:@127.0.0.1/qmoon_aquarius-2001?sslmode=disable
		return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dc.User, dc.Password, dc.Addr, dc.Database)
	default:
		return ""
	}
}

// DefaultRPCConfig returns a default configuration for the HttpServer server
func DefaultDBConfig() *DBConfig {
	return &DBConfig{
		DriverName: "postgres",
		User:       "postgres",
		Password:   "Postgres@",
		Addr:       "localhost:5432",
		Database:   "qmoon",
	}
}

// TestRPCConfig returns a configuration for testing the HttpServer server
func TestDBConfig() *DBConfig {
	return &DBConfig{
		DriverName: "postgres",
		User:       "postgres",
		Password:   "Postgres@",
		Addr:       "localhost:5432",
		Database:   "qmoon",
	}
}
