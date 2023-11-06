package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

const defaultConfigFilePath string = "/etc/goboy_config.yml"

// Config is a struct that will hold all GoBoy configuration
type Config struct {
	RomPath         string `yaml:"rom_path"`
	LogStdoutEnable bool   `yaml:"log_stdout_enable"`
	LogTruncate     bool   `yaml:"log_truncate"`
	LogFilePath     string `yaml:"log_file_path"`
}

func (c *Config) checkEssentialValues() (bool, string) {
	var missingValues []string
	if len(c.RomPath) == 0 {
		missingValues = append(missingValues, "rom_path")
	}

	if len(c.LogFilePath) == 0 {
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
	info, err := os.Stat(c.LogFilePath)
	if err != nil {
		return false
	}

	return info.Mode().Perm()&os.FileMode(0200) != 0
}

func NewConfigParser(configPath string) (*Parser, error) {
	if len(configPath) == 0 {
		configPath = defaultConfigFilePath
	}

	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		return nil, ConfigurationFileNotFound{err.Error()}
	}

	return &Parser{defaultConfigFilePath}, nil
}

func (p Parser) Parse() (*Config, error) {
	bytes, err := os.ReadFile(p.configFilePath)
	if err != nil { // this error might be occurring because we do not have permissions to read it
		return nil, err
	}

	var configStruct *Config
	err = yaml.Unmarshal(bytes, configStruct)
	if err != nil {
		return nil, ParsingError{err.Error()}
	}

	checkConfigParameters, missingParameters := configStruct.checkEssentialValues()
	if !checkConfigParameters {
		return nil, MissingConfigValues{missingParameters}
	}

	if !configStruct.checkRomExist() {
		return nil, RomNotFoundError{configStruct.RomPath}
	}

	if !configStruct.checkLogPathIsWritable() {
		return nil, LogWriteError{configStruct.LogFilePath}
	}

	return configStruct, nil
}
