package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/arsmoriendy/opor/gql-srv/internal"
	"github.com/arsmoriendy/opor/gql-srv/internal/loglvl"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	pgxuuid "github.com/vgarvardt/pgx-google-uuid/v5"
)

// Make sure to initialize using InitPool().
// Don't forget to close.
var Pool *pgxpool.Pool

func InitPool() {
	dbcfg, err := pgxpool.ParseConfig(os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}
	dbcfg.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxuuid.Register(conn.TypeMap())
		return nil
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), dbcfg)
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
	row := Pool.QueryRow(context.Background(), "select expire_at from uuids where uuid = $1", uuid)

	var expire *time.Time
	err = row.Scan(&expire)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return false, nil
		}
		return
	}
	if expire != nil {
		return expire.After(time.Now()), nil
	}
	return true, nil
}

func NewUuid() (uuidStr string, err error) {
	row := Pool.QueryRow(context.Background(), "insert into uuids default values returning uuid")

	var uuid uuid.UUID
	err = row.Scan(&uuid)
	if err != nil {
		return
	}

	uuidJson, err := uuid.MarshalText()
	if err != nil {
		return
	}

	return string(uuidJson), nil
}

func NewFrontUuid() (uuidStr string, err error) {
	row := Pool.QueryRow(context.Background(),
		"insert into uuids(expire_at) values($1) returning uuid",
		time.Now().Add(internal.FrontUuidExpr))

	var uuid uuid.UUID
	err = row.Scan(&uuid)
	if err != nil {
		return
	}

	uuidJson, err := uuid.MarshalText()
	if err != nil {
		return
	}

	return string(uuidJson), nil
}

func RmUuid(uuidStr string) (err error) {
	uuid, err := uuid.Parse(uuidStr)
	if err != nil {
		return
	}

	_, err = Pool.Exec(context.Background(), "delete from uuids where uuid = $1", uuid)
	if err != nil {
		return
	}

	return
}
