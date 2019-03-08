package log

import (
	"os"
	"testing"
)

// default level should be Info
func TestLevelDefault(t *testing.T) {
	l := GetLevel()
	if l != Info {
		t.Errorf("expected Info, got %s", l)
	}

	os.Setenv("LOG_LEVEL", "Debug")
	l = Default()
	if l != Debug {
		t.Errorf("expected Debug, got %s", l)
	}
	os.Unsetenv("LOG_LEVEL")
}

// set the level and see that it sticks
func TestLevelSet(t *testing.T) {
	SetLevel(None)
	l := GetLevel()
	if l != None {
		t.Errorf("expected None, got %s", l)
	}
}

func TestLevelGet(t *testing.T) {
	// try to set a bogus level
	l := GetLevel()
	SetLevel(42)
	n := GetLevel()
	if n != l {
		t.Errorf("expected %s, got %s", l, n)
	}
}

func TestAllLevels(t *testing.T) {
	// set a specific level and gets its string value
	data := map[Level]string{
		None:  "None",
		Trace: "Trace",
		Debug: "Debug",
		Info:  "Info",
		Warn:  "Warn",
		Error: "Error",
	}

	for l, e := range data {
		SetLevel(l)
		s := GetLevel().String()
		if s != e {
			t.Errorf("expected logging level %s, got %s", e, s)
		}
	}
}

// try to get a bogus level string
func TestLevelUnknown(t *testing.T) {
	s := Level(23).String()
	if s != "unknown" {
		t.Errorf("expected unknown, got %s", s)
	}
}
