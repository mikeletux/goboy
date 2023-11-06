package config

import (
	"fmt"
	"path/filepath"
)

// ConfigurationFileNotFound is an error intended to be return when GoBoy configuration file does not exist
type ConfigurationFileNotFound struct {
	filePath string
}

func (c *ConfigurationFileNotFound) Error() string {
	if len(c.filePath) == 0 {
		return "Configuration file not specified."
	}
	return fmt.Sprintf("Configuration file not set and %s does not exist\n", c.filePath)
}

// ParsingError is an error intended to be returned when having issues parsing the config file
type ParsingError struct {
	filePath string
}

func (p *ParsingError) Error() string {
	return fmt.Sprintf("Error parsing file %s. Please check syntax.", p.filePath)
}

// MissingConfigValuesError is an error intended to be returned when not all parameters are set in the config file
type MissingConfigValuesError struct {
	missingParameters string
}

func (m *MissingConfigValuesError) Error() string {
	return fmt.Sprintf("The following parameters are missing or empty in the config file: %s",
		m.missingParameters)
}

// RomNotFoundError is an error intended to be returned when the rom specified does not exist
type RomNotFoundError struct {
	romPath string
}

func (r *RomNotFoundError) Error() string {
	return fmt.Sprintf("Rom file %s does not exist", r.romPath)
}

// LogWriteError is an error intended to be returned when the user does not have permission to write in the
// location specified
type LogWriteError struct {
	filePath string
}

func (l *LogWriteError) Error() string {
	return fmt.Sprintf("Folder %s does not exist or do not have permission to write the log.",
		filepath.Dir(l.filePath))
}
