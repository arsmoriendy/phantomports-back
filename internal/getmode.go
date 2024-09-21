package internal

import "os"

func GetMode() (string, bool) {
	return os.LookupEnv("MODE")
}

func IsDevMode() bool {
	mode, found := GetMode()
	if !found {
		return true
	}
	if mode == "PROD" {
		return false
	}
	return true
}
