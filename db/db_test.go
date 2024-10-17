package db

import (
	"os"
	"testing"
	"time"

	"github.com/arsmoriendy/opor/gql-srv/internal"
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

	InitPool()
	defer Pool.Close()

	uuid, err := NewUuid()
	if err != nil {
		t.Fatal(err)
	}

	valid, err := UuidValid(uuid)
	if err != nil {
		t.Fatal(err)
	}

	if !valid {
		t.Fatal(valid)
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

	InitPool()
	defer Pool.Close()

	uuid, err := NewFrontUuid()
	if err != nil {
		t.Fatal(err)
	}

	valid, err := UuidValid(uuid)
	if err != nil {
		t.Fatal(err)
	}

	if !valid {
		t.Fatal(valid)
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

	valid, err := UuidValid(uuid)
	if err != nil {
		t.Fatal(err)
	}

	if valid {
		t.Fatal(valid)
	}

	err = RmUuid(uuid)
	if err != nil {
		t.Fatal(err)
	}
}
