package internal

import (
	"errors"
	"log"
	"os"

	sll "github.com/arsmoriendy/sixlvllog"
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

	sll.Init()

	FrontUuidLifetime = getFrontUuidLifetime()
	if sll.LogLvl >= sll.INFO {
		log.Printf("set front uuid lifetime to: %s", FrontUuidLifetime)
	}

	RefInterval = getRefInterval()
	if sll.LogLvl >= sll.INFO {
		log.Printf("refreshing ports every %s", RefInterval)
	}
}
