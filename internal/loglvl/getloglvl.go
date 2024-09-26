package loglvl

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type LogLvlT uint

// Memoized runtime LogLvlT
var LogLvl LogLvlT

const (
	FATAL = iota
	ERROR
	WARN
	INFO
	DEBUG
	TRACE
)

const DEFAULT_LOG_LEVEL = INFO

// LogLvl to string
var lvlStr = map[LogLvlT]string{
	FATAL: "FATAL",
	ERROR: "ERROR",
	WARN:  "WARN",
	INFO:  "INFO",
	DEBUG: "DEBUG",
	TRACE: "TRACE",
}

var ErrInvalidLvlStr = errors.New("invalid log level string")

// string to LogLvl
func checkLvlStr(lvl string) (LogLvlT, error) {
	for k, v := range lvlStr {
		if strings.EqualFold(lvl, v) {
			return k, nil
		}
	}
	return 0, fmt.Errorf("%w: %s", ErrInvalidLvlStr, lvl)
}

var ErrOutOfBounds = errors.New("log level out of bounds")

func checkLvlBound(lvl LogLvlT) error {
	if lvl < FATAL || lvl > TRACE {
		return fmt.Errorf("%w: %v", ErrOutOfBounds, lvl)
	}
	return nil
}

// Use memoized var LogLvl instead, but don't forget to call Init().
func Get() LogLvlT {
	lvlenv := os.Getenv("LOG_LEVEL")

	lvlNum, err := strconv.Atoi(lvlenv)
	if err != nil {
		goto CHECK_STR
	}

	// assume $LOG_LEVEL a number
	err = checkLvlBound(LogLvlT(lvlNum))
	if err != nil {
		log.Println(
			fmt.Errorf("%w: using default level: %v\n", err, lvlStr[DEFAULT_LOG_LEVEL]))
		return DEFAULT_LOG_LEVEL
	}
	return LogLvlT(lvlNum)

CHECK_STR:
	// assume $LOG_LEVEL is a stirng
	lvlstring, err := checkLvlStr(lvlenv)
	if err != nil {
		return DEFAULT_LOG_LEVEL
	}

	return lvlstring
}

func Init() {
	LogLvl = Get()
}
