package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/arsmoriendy/opor/gql-srv/internal/loglvl"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Make sure to initialize using InitPool().
// Don't forget to close.
var Pool *pgxpool.Pool

func InitPool() {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		pool.Close()
		panic(fmt.Errorf("unable to create connection pool: %w", err))
	}
	if loglvl.LogLvl >= loglvl.INFO {
		log.Printf("connected to database")
	}
	Pool = pool
}

func UuidExists(uuid string) (exists bool, err error) {
	row := Pool.QueryRow(context.Background(), "select 1 from uuids where uuid = $1", uuid)

	var existsInt int
	err = row.Scan(&existsInt)
	if err.Error() == "no rows in result set" {
		return false, nil
	}
	if err != nil {
		return
	}
	return existsInt == 1, nil
}
