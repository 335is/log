package log

import (
	"os"
	"strings"
)

// Level defines the logging level type
type Level int

// Log level controls the logging threshold, which log statements will be output.
// For example, Debug level will not output Trace statements, but will output Debug, Info, and the rest.
const (
	// None is no logging
	None Level = iota

	// Trace logs verbose details and, really should be metricss
	Trace

	// Debug logs Debug and above
	Debug

	// Info logs Info and above
	Info

	// Warn logs Warn and above
	Warn

	// Error logs Error and above
	Error
)

var (
	level  = Default()
	levels = []string{"None", "Trace", "Debug", "Info", "Warn", "Error"}
)

// Default returns the default log level
// LOG_LEVEL env var overrides the default of Info
func Default() Level {
	s := os.Getenv("LOG_LEVEL")
	l := FromString(s)
	if l == -1 {
		return Info
	}

	return l
}

// String returns the level as a string
func (l Level) String() string {
	if l >= None && l <= Error {
		return levels[l]
	}

	return "unknown"
}

// FromString converts the string representation of level into its const
// It does a case-insensitive comparison
// Invalid level string returns -1
func FromString(s string) Level {
	for l, str := range levels {
		if strings.EqualFold(s, str) {
			return Level(l)
		}
	}

	return -1
}

// GetLevel returns the current logging level
func GetLevel() Level {
	return level
}

// SetLevel sets the current logging level
func SetLevel(l Level) {
	if l >= None && l <= Error {
		level = l
	}
}
