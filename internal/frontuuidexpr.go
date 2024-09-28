package internal

import (
	"os"
	"strconv"
	"time"
	"github.com/joho/godotenv"
)

var FrontUuidExpr = getFrontUuidExpr()

func getFrontUuidExpr() (expr time.Duration) {
	godotenv.Load()
	expr_env := os.Getenv("FRONT_UUID_EXPR")
	if expr_int, err := strconv.Atoi(expr_env); err != nil {
		expr = time.Hour
	} else {
		expr = time.Duration(expr_int) * time.Millisecond
	}
	return
}

func ResetFrontUuidExpr() {
	FrontUuidExpr = getFrontUuidExpr()
}
