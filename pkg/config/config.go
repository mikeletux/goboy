package config

import (
	"golang.org/x/sys/unix"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

const defaultConfigFilePath string = "/etc/goboy_config.yml"

// Config is a struct that will hold all GoBoy configuration
type Config struct {
	RomPath         string `yaml:"rom_path"`
	LogStdoutEnable bool   `yaml:"log_stdout_enable"`
	LogFileEnable   bool   `yaml:"log_file_enable"`
	LogFilePath     string `yaml:"log_file_path"`
}

func (c *Config) checkEssentialValues() (bool, string) {
	var missingValues []string
	if len(c.RomPath) == 0 {
		missingValues = append(missingValues, "rom_path")
	}

	if c.LogFileEnable && len(c.LogFilePath) == 0 {
		missingValues = append(missingValues, "log_file_path")
	}

	if len(missingValues) > 0 {
		return false, strings.Join(missingValues, ",")
	}

	return true, ""
}

func (c *Config) checkRomExist() bool {
	_, err := os.Stat(c.RomPath)
	if os.IsNotExist(err) {
		return false
	}

	return true
}

type Parser struct {
	configFilePath string
}

func (c *Config) checkLogPathIsWritable() bool {
	_, err := os.Stat(filepath.Dir(c.LogFilePath))
	if err != nil {
		return false
	}

	return unix.Access(filepath.Dir(c.LogFilePath), unix.W_OK) == nil
}

func NewConfigParser(configPathIn string) (*Parser, error) {
	configPath := configPathIn
	if len(configPath) == 0 {
		configPath = defaultConfigFilePath
	}

	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		return nil, &ConfigurationFileNotFound{configPath}
	}

	return &Parser{configPath}, nil
}

func (p Parser) Parse() (*Config, error) {
	bytes, err := os.ReadFile(p.configFilePath)
	if err != nil { // this error might be occurring because we do not have permissions to read it
		return nil, err
	}

	configStruct := &Config{}
	err = yaml.Unmarshal(bytes, configStruct)
	if err != nil {
		return nil, &ParsingError{p.configFilePath}
	}

	checkConfigParameters, missingParameters := configStruct.checkEssentialValues()
	if !checkConfigParameters {
		return nil, &MissingConfigValuesError{missingParameters}
	}

	if !configStruct.checkRomExist() {
		return nil, &RomNotFoundError{configStruct.RomPath}
	}

	if configStruct.LogFileEnable && !configStruct.checkLogPathIsWritable() {
		return nil, &LogWriteError{configStruct.LogFilePath}
	}

	return configStruct, nil
}
