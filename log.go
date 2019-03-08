package log

import (
	"io"
	golog "log"
	"os"
	"strings"
)

var (
	logger = golog.New(os.Stderr, "", golog.Ldate|golog.Lmicroseconds)
)

// Tracef logs Trace level and above
func Tracef(format string, vars ...interface{}) {
	if level > None && level <= Trace {
		logger.Printf(strings.Join([]string{"TRACE", format}, " "), vars...)
	}
}

// Debugf logs Debug level and above
func Debugf(format string, vars ...interface{}) {
	if level > None && level <= Debug {
		logger.Printf(strings.Join([]string{"DEBUG", format}, " "), vars...)
	}
}

// Infof logs Info level and above
func Infof(format string, vars ...interface{}) {
	if level > None && level <= Info {
		logger.Printf(strings.Join([]string{"INFO", format}, " "), vars...)
	}
}

// Warnf logs Warn level and above
func Warnf(format string, vars ...interface{}) {
	if level > None && level <= Warn {
		logger.Printf(strings.Join([]string{"WARN", format}, " "), vars...)
	}
}

// Errorf logs Error level and above
func Errorf(format string, vars ...interface{}) {
	if level > None && level <= Error {
		logger.Printf(strings.Join([]string{"ERROR", format}, " "), vars...)
	}
}

// SetWriter sets the writer that our logger uses, such as os.Stderr, os.Stdout, a file writer, etc.
func SetWriter(w io.Writer) {
	logger.SetOutput(w)
}
