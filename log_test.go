package log

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

// level None should not output anything
func TestLoggerNone(t *testing.T) {
	var b bytes.Buffer
	SetWriter(&b)
	SetLevel(None)
	Infof("%s", "watch this message disappear")

	s := b.String()
	if len(s) != 0 {
		t.Errorf("expected no log output, got %s", s)
	}
}

// level Trace only outputs Tracef messages
func TestLoggerTrace(t *testing.T) {
	var b bytes.Buffer
	SetWriter(&b)
	SetLevel(Trace)
	msg := "trace message"
	Tracef("%s", msg)

	s := b.String()
	if !strings.Contains(s, msg) {
		t.Errorf("expected log output to contain %s, got %s", msg, s)
	}

	b.Reset()
	SetLevel(Debug)
	Tracef("%s", msg)
	s = b.String()
	if len(s) != 0 {
		t.Errorf("expected no log output, got %s", s)
	}
}

// basic test of log Debug
func TestLoggerDebug(t *testing.T) {
	var b bytes.Buffer
	SetWriter(&b)
	SetLevel(Debug)
	msg := "debug message"
	Debugf("%s", msg)

	s := b.String()
	if !strings.Contains(s, msg) {
		t.Errorf("expected log string to contain %s, got %s", msg, s)
	}

	b.Reset()
	SetLevel(Info)
	Debugf("%s", msg)
	s = b.String()
	if len(s) != 0 {
		t.Errorf("expected no log output, got %s", s)
	}
}

// basic test of log Info
func TestLoggerInfo(t *testing.T) {
	var b bytes.Buffer
	SetWriter(&b)
	SetLevel(Info)
	msg := "info message"
	Infof("%s", msg)

	s := b.String()
	if !strings.Contains(s, msg) {
		t.Errorf("expected log string to contain %s, got %s", msg, s)
	}

	b.Reset()
	SetLevel(Warn)
	Infof("%s", msg)
	s = b.String()
	if len(s) != 0 {
		t.Errorf("expected no log output, got %s", s)
	}
}

// basic test of log Warn
func TestLoggerWarn(t *testing.T) {
	var b bytes.Buffer
	SetWriter(&b)
	SetLevel(Warn)
	msg := "warn message"
	Warnf("%s", msg)

	s := b.String()
	if !strings.Contains(s, msg) {
		t.Errorf("expected log string to contain %s, got %s", msg, s)
	}

	b.Reset()
	SetLevel(Error)
	Warnf("%s", msg)
	s = b.String()
	if len(s) != 0 {
		t.Errorf("expected no log output, got %s", s)
	}
}

// basic test of log Error
func TestLoggerError(t *testing.T) {
	var b bytes.Buffer
	SetWriter(&b)
	SetLevel(Error)
	msg := "error message"
	Errorf("%s", msg)

	s := b.String()
	if !strings.Contains(s, msg) {
		t.Errorf("expected log string to contain %s, got %s", msg, s)
	}
}

func TestLogToFile(t *testing.T) {
	file, err := ioutil.TempFile(".", "logtest_")
	if err != nil {
		t.Fatalf("%v", err)
	}

	SetWriter(file)
	SetLevel(Error)
	msg := "error message in a file"
	Errorf("%s", msg)

	file.Close()

	data, err := ioutil.ReadFile(file.Name())
	if err != nil {
		t.Fatalf("%v", err)
	}

	s := string(data)
	if !strings.Contains(s, msg) {
		t.Errorf("expected log file to contain %s, got %s", msg, s)
	}

	os.Remove(file.Name())
}
