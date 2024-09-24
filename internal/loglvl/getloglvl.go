package loglvl

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type LogLvl uint

const (
	FATAL = iota
	ERROR
	WARN
	INFO
	DEBUG
	TRACE
)

// LogLvl to string
var lvlStr = map[LogLvl]string{
	FATAL: "FATAL",
	ERROR: "ERROR",
	WARN:  "WARN",
	INFO:  "INFO",
	DEBUG: "DEBUG",
	TRACE: "TRACE",
}

var ErrInvalidLvlStr = errors.New("invalid log level string")

// string to LogLvl
func checkLvlStr(lvl string) (LogLvl, error) {
	for k, v := range lvlStr {
		if strings.EqualFold(lvl, v) {
			return k, nil
		}
	}
	return 0, fmt.Errorf("%w: %s", ErrInvalidLvlStr, lvl)
}

var ErrOutOfBounds = errors.New("log level out of bounds")

func checkLvlBound(lvl LogLvl) error {
	if lvl < FATAL || lvl > TRACE {
		return fmt.Errorf("%w: %v", ErrOutOfBounds, lvl)
	}
	return nil
}

func Get() (LogLvl, error) {
	lvlStr := os.Getenv("LOG_LEVEL")

	lvlNum, err := strconv.Atoi(lvlStr)
	if err != nil {
		goto CHECK_STR
	}

	// assume $LOG_LEVEL a number
	err = checkLvlBound(LogLvl(lvlNum))
	if err != nil {
		return 0, err
	}
	return LogLvl(lvlNum), nil

CHECK_STR:
	// assume $LOG_LEVEL is a stirng
	lvlstring, err := checkLvlStr(lvlStr)
	if err != nil {
		return 0, err
	}

	return lvlstring, nil
}
