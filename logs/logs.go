package logs

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Logging Level.
const (
	Off = iota
	Trace
	Debug
	Info
	Warn
	Error
	Fatal
)

// all loggers.
var loggers []*Logger

var logLevel = Debug

// Logger represents a simple logger with level.
type Logger struct {
	level  int
	logger *log.Logger
}

// NewLogger creates a logger
func NewLogger(out io.Writer) *Logger {
	ret := &Logger{level: logLevel, logger: log.New(out, "", log.LstdFlags|log.Lshortfile)}

	loggers = append(loggers, ret)
	return ret
}

// SetLevel sets the logging level of all loggers
func SetLevel(level string) {
	logLevel = getLevel(level)
	for _, l := range loggers {
		l.SetLevel(level)
	}
}

func getLevel(level string) int {
	level = strings.ToLower(level)
	switch level {
	case "off":
		return Off
	case "trace":
		return Trace
	case "debug":
		return Debug
	case "info":
		return Info
	case "warn":
		return Warn
	case "error":
		return Error
	case "fatal":
		return Fatal
	default:
		return Info
	}
}

// SetLevel sets the logging level of a logger.
func (l *Logger) SetLevel(level string) {
	l.level = getLevel(level)
}

// IsTraceEnabled determines whether the trace level is enabled.
func (l *Logger) IsTraceEnabled() bool {
	return l.level <= Trace
}

// IsDebugEnabled determines whether the debug level is enabled.
func (l *Logger) IsDebugEnabled() bool {
	return l.level <= Debug
}

// IsWarnEnabled determines whether the debug level is enabled.
func (l *Logger) IsWarnEnabled() bool {
	return l.level <= Warn
}

// Trace prints trace level message
func (l *Logger) Trace(v ...interface{}) {
	if Trace < l.level {
		return
	}
	l.logger.SetPrefix("Trace ")
	l.logger.Output(2, fmt.Sprint(v...))
}

// Tracef prints trace level message with format.
func (l *Logger) Tracef(format string, v ...interface{}) {
	if Trace < l.level {
		return
	}

	l.logger.SetPrefix("Trace ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}

// Debug prints debug level message.
func (l *Logger) Debug(v ...interface{}) {
	if Debug < l.level {
		return
	}

	l.logger.SetPrefix("Debug ")
	l.logger.Output(2, fmt.Sprint(v...))
}

// Debugf prints debug level message with format.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if Debug < l.level {
		return
	}

	l.logger.SetPrefix("Debug ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}

// Info prints info level message.
func (l *Logger) Info(v ...interface{}) {
	if Info < l.level {
		return
	}
	l.logger.SetPrefix("Info ")
	l.logger.Output(2, fmt.Sprint(v...))
}

// Infof prints info level message with format.
func (l *Logger) Infof(format string, v ...interface{}) {
	if Info < l.level {
		return
	}

	l.logger.SetPrefix("Info ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}

// Warn prints warning level message.
func (l *Logger) Warn(v ...interface{}) {
	if Warn < l.level {
		return
	}

	l.logger.SetPrefix("Warn ")
	l.logger.Output(2, fmt.Sprint(v...))
}

// Warn prints warning level message with format.
func (l *Logger) Warnf(format string, v ...interface{}) {
	if Warn < l.level {
		return
	}

	l.logger.SetPrefix("Warn ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}

// Error prints error level message.
func (l *Logger) Error(v ...interface{}) {
	if Error < l.level {
		return
	}

	l.logger.SetPrefix("Error ")
	l.logger.Output(2, fmt.Sprint(v...))
}

// Errorf prints error level message with format.
func (l *Logger) Errorf(format string, v ...interface{}) {
	if Error < l.level {
		return
	}

	l.logger.SetPrefix("Error ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}

// Fatal prints fatal level message.
func (l *Logger) Fatal(v ...interface{}) {
	if Fatal < l.level {
		return
	}

	l.logger.SetPrefix("Fatal ")
	l.logger.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf prints fatal level message with format.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if Fatal < l.level {
		return
	}

	l.logger.SetPrefix("Fatal ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}
