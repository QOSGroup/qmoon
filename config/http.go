// Copyright 2018 The QOS Authors

package config

// HttpServerConfig defines the configuration options for the HttpServer server
type HttpServerConfig struct {
	// TCP or UNIX socket address for the HttpServer server to listen on
	ListenAddress string `mapstructure:"laddr"`
}

// DefaultRPCConfig returns a default configuration for the HttpServer server
func DefaultRPCConfig() *HttpServerConfig {
	return &HttpServerConfig{
		ListenAddress: "http://0.0.0.0:8080",
	}
}

// TestRPCConfig returns a configuration for testing the HttpServer server
func TestRPCConfig() *HttpServerConfig {
	cfg := DefaultRPCConfig()
	cfg.ListenAddress = "http://0.0.0.0:18080"
	return cfg
}
