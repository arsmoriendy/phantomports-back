package db

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/arsmoriendy/opor/gql-srv/internal"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func TestNewUuid(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	InitPool()
	defer Pool.Close()

	uuid, err := NewUuid()
	if err != nil {
		t.Fatal(err)
	}

	err = RmUuid(uuid)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewFrontUuid(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	InitPool()
	defer Pool.Close()

	uuid, err := NewFrontUuid()
	if err != nil {
		t.Fatal(err)
	}

	err = RmUuid(uuid)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUuidValid(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
	internal.ResetFrontUuidLifetime()

	InitPool()
	defer Pool.Close()

	uuid, err := NewUuid()
	if err != nil {
		t.Fatal(err)
	}

	err = UuidValid(uuid)
	if err != nil {
		t.Fatal(err)
	}

	err = RmUuid(uuid)
	if err != nil {
		t.Fatal(err)
	}
}

func TestValidFrontUuid(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
	internal.ResetFrontUuidLifetime()

	InitPool()
	defer Pool.Close()

	uuid, err := NewFrontUuid()
	if err != nil {
		t.Fatal(err)
	}

	err = UuidValid(uuid)
	if err != nil {
		t.Fatal(err)
	}

	err = RmUuid(uuid)
	if err != nil {
		t.Fatal(err)
	}
}

func TestExpiredFrontUuid(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	os.Setenv("FRONT_UUID_EXPR", "1")
	internal.ResetFrontUuidLifetime()

	InitPool()
	defer Pool.Close()

	uuid, err := NewFrontUuid()
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(internal.FrontUuidLifetime)

	err = UuidValid(uuid)
	if err == nil {
		t.Fail()
	}
	if err != nil && !errors.Is(err, ErrExpiredUuid) {
		t.Fatal(err)
	}

	err = RmUuid(uuid)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnregisteredUuid(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	InitPool()
	defer Pool.Close()

	uuid := uuid.New().String()

	err = UuidValid(uuid)
	if err == nil {
		t.Fail()
	}

	if !errors.Is(err, ErrUnregisteredUuid) {
		t.Fatal(err)
	}

	err = RmUuid(uuid)
	if err != nil {
		t.Fatal(err)
	}
}
