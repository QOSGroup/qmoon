// Copyright 2018 The QOS Authors

package config

import (
	"bytes"
	"html/template"
	"path/filepath"
	"fmt"
	cmn "github.com/tendermint/tendermint/libs/common"
)

var configTemplate *template.Template

func init() {
	var err error
	if configTemplate, err = template.New("configFileTemplate").Parse(defaultConfigTemplate); err != nil {
		panic(err)
	}
}

// WriteDefaultConfigFile 初始化默认配置文件
func WriteDefaultConfigFile(rootDir string) {
	if err := cmn.EnsureDir(rootDir, 0700); err != nil {
		panic(fmt.Sprintf("Panicked on a Crisis: %v", err.Error()))
	}

	if err := cmn.EnsureDir(filepath.Join(rootDir, defaultConfigDir), 0700); err != nil {
		panic(fmt.Sprintf("Panicked on a Crisis: %v", err.Error()))
	}

	configFilePath := filepath.Join(rootDir, defaultConfigFilePath)

	if !cmn.FileExists(configFilePath) {
		WriteConfigFile(configFilePath, DefaultConfig())
	}
}

// WriteConfigFile renders config using the template and writes it to configFilePath.
func WriteConfigFile(configFilePath string, config *Config) {
	var buffer bytes.Buffer

	if err := configTemplate.Execute(&buffer, config); err != nil {
		panic(err)
	}

	cmn.MustWriteFile(configFilePath, buffer.Bytes(), 0644)
}

const defaultConfigTemplate = `# This is a TOML config file.
# For more information, see https://github.com/toml-lang/toml

##### main base config options #####

# qmoon directory
home = "{{ .BaseConfig.RootDir }}"

# Output level for logging, including package level options
log_level = "{{ .BaseConfig.LogLevel }}"

##### http server configuration options #####
[http]
# Database type
laddr = "{{ .HttpServer.ListenAddress }}"

##### db configuration options #####
[db]

# Database type
driver = "{{ .DB.DriverName }}"

# Database user
user = "{{ .DB.User }}"

# Database password
password = "{{ .DB.Password }}"

# Database addr
addr = "{{ .DB.Addr }}"

# Database database
database = "{{ .DB.Database }}"
`
