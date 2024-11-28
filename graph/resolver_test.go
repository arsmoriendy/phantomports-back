package graph

import (
	"encoding/csv"
	"errors"
	"os"
	"strings"
	"testing"
)

var r = Resolver{}

func TestFillPorts(t *testing.T) {
	f, err := os.Open("../test/data/service-names-port-numbers.csv")
	defer f.Close()
	if err != nil {
		t.Fatal("cannot find mock csv file, check pathing")
	}

	rdr := csv.NewReader(f)

	err = r.fillPorts(rdr)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFillPortsEmpty(t *testing.T) {
	rdr := csv.NewReader(strings.NewReader(""))

	err := r.fillPorts(rdr)
	if !errors.Is(err, ErrEmptyPortCsv) {
		t.Fatal("unexpected error")
	}
}
