package db

import (
	"testing"

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
