package config

// ConfigurationFileNotFound is an error intended to be return when GoBoy configuration file does not exist
type ConfigurationFileNotFound struct {
	message string
}

func (c ConfigurationFileNotFound) Error() string {
	return c.message
}

// ParsingError is an error intended to be returned when having issues parsing the config file
type ParsingError struct {
	message string
}

func (p ParsingError) Error() string {
	return p.message
}

// RomNotFoundError is an error intended to be returned when the rom specified does not exist
type RomNotFoundError struct {
	message string
}

func (r RomNotFoundError) Error() string {
	return r.message
}

// LogWriteError is an error intended to be returned when the user does not have permission to write in the
// location specified
type LogWriteError struct {
	message string
}

func (l LogWriteError) Error() string {
	return l.message
}

// MissingConfigValues is an error intended to be returned when not all parameters are set in the config file
type MissingConfigValues struct {
	message string
}

func (m MissingConfigValues) Error() string {
	return m.message
}
