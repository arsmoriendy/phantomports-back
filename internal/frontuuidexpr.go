package internal

import (
	"os"
	"strconv"
	"time"
)

var FrontUuidExpr = getRefInterval()

func getFrontUuidExpr() (expr time.Duration) {
	expr_env := os.Getenv("FRONT_UUID_EXPR")
	if expr_int, err := strconv.Atoi(expr_env); err != nil {
		expr = time.Hour
	} else {
		expr = time.Duration(expr_int) * time.Millisecond
	}
	return
}