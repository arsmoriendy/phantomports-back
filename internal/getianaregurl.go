package internal

import "os"

var IANAregUrl string

func getIANAregUrl() string {
	env := os.Getenv("IANA_REG_URL")
	if env == "" {
		return "https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.csv"
	}
	return env
}
