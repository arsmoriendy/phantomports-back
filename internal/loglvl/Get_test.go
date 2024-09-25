package loglvl

import (
	"os"
	"testing"
)

// Test Get on case insensitive string
func TestStr(t *testing.T) {
	err := os.Setenv("LOG_LEVEL", "TrAcE")
	if err != nil {
		t.Fatal(err)
	}

	lvl := Get()
	if lvl != TRACE {
		t.Fatal(lvl)
	}
}

// Number string
func TestNumStr(t *testing.T) {
	err := os.Setenv("LOG_LEVEL", "5")
	if err != nil {
		t.Fatal(err)
	}

	lvl := Get()
	if lvl != TRACE {
		t.Fatal(lvl)
	}
}

// NaN
func TestNan(t *testing.T) {
	err := os.Setenv("LOG_LEVEL", "NaN")
	if err != nil {
		t.Fatal(err)
	}

	lvl := Get()
	if lvl != INFO {
		t.Fatal(lvl)
	}
}

func TestOverbound(t *testing.T) {
	err := os.Setenv("LOG_LEVEL", "6")
	if err != nil {
		t.Fatal(err)
	}

	lvl := Get()
	if lvl != INFO {
		t.Fatal(lvl)
	}
}
