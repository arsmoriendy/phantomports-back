package internal

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var RefInterval = getRefInterval()

func getRefInterval() (ri time.Duration) {
	godotenv.Load()
	ri_env := os.Getenv("REFRESH_INTERVAL")
	if ri_int, err := strconv.Atoi(ri_env); err != nil {
		ri = time.Hour
	} else {
		ri = time.Duration(ri_int) * time.Millisecond
	}
	return
}
