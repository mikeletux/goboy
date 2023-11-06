package log

import (
	"io"
	"log"
	"os"
)

// Logger is the interface used by other emulator components to log.
type Logger interface {
	Debug(args ...any)
	Debugf(format string, args ...any)
	Fatal(args ...any)
	Fatalf(format string, args ...any)
	Close()
}

// BuiltinLogger is a simple wrapper to log.Logger.
type BuiltinLogger struct {
	noLog   bool
	logFile *os.File
	logger  *log.Logger
}

// NewBuiltinStdoutLogger returns a simple BuiltinLogger that only logs on stdout.
func NewBuiltinStdoutLogger(logStdout bool, logFileEnable bool, logPath string) (*BuiltinLogger, error) {
	var noLog bool
	var logFile *os.File
	var writer io.Writer

	switch {
	case !logStdout && !logFileEnable:
		noLog = true
		writer = os.Stdout
		break

	case logStdout && logFileEnable && len(logPath) > 0:
		logFile, err := openFile(logPath)
		if err != nil {
			return nil, err
		}
		writer = io.MultiWriter(os.Stdout, logFile)
		break

	case logStdout:
		writer = os.Stdout
		break

	case logFileEnable && len(logPath) > 0:
		logFile, err := openFile(logPath)
		if err != nil {
			return nil, err
		}
		writer = logFile
		break
	}

	return &BuiltinLogger{
		noLog:   noLog,
		logFile: logFile,
		logger:  log.New(writer, "", 0),
	}, nil
}

func (b *BuiltinLogger) Debug(args ...any) {
	if !b.noLog {
		b.logger.Println(args...)
	}
}
func (b *BuiltinLogger) Debugf(format string, args ...any) {
	if !b.noLog {
		b.logger.Printf(format, args...)
	}
}
func (b *BuiltinLogger) Fatal(args ...any) {
	b.logger.Fatalln(args...)
}
func (b *BuiltinLogger) Fatalf(format string, args ...any) {
	b.logger.Fatalf(format, args...)
}

func (b *BuiltinLogger) Close() {
	if b.logFile != nil {
		b.logFile.Close()
	}
}

func openFile(logFile string) (*os.File, error) {
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return nil, err
	}

	return f, nil
}
