package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/arsmoriendy/opor/gql-srv/internal"
	sll "github.com/arsmoriendy/sixlvllog"
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
	if sll.LogLvl >= sll.INFO {
		log.Printf("connected to database")
	}

	Pool = pool
}

var ErrUnregisteredUuid = errors.New("unregistered uuid")
var ErrExpiredUuid = errors.New("Expired uuid")

func UuidValid(uuidStr string) (err error) {
	uuid_uuid, err := uuid.Parse(uuidStr)
	if err != nil {
		return
	}

	row := Pool.QueryRow(context.Background(), "select expire_at from uuids where uuid = $1", uuid_uuid)

	var expire *time.Time
	err = row.Scan(&expire)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("%w: %s", ErrUnregisteredUuid, uuidStr)
		}
		return
	}
	if expire != nil && expire.Before(time.Now()) {
		return fmt.Errorf("%w: %s", ErrExpiredUuid, uuidStr)
	}
	return
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
		time.Now().Add(internal.FrontUuidLifetime))

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
