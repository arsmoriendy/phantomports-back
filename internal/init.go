package internal

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	if IsDevMode() {
		log.SetPrefix("[DEV] ")
	} else {
		log.SetPrefix("[PROD] ")
	}

	// load env vars
	err := godotenv.Load()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		panic(err)
	}
}
