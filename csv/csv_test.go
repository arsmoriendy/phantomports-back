package csv

import (
	"errors"
	"net/http"
	"testing"
)

func TestValidateResInvalidStatus(t *testing.T) {
	res := http.Response{
		StatusCode: http.StatusNotFound,
		Header: map[string][]string{
			"Content-Type": {"text/csv"},
		},
	}

	err := validateRes(&res)
	if !errors.Is(err, ErrInvalidStatusCode) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateResType(t *testing.T) {
	res := http.Response{
		StatusCode: http.StatusOK,
		Header: map[string][]string{
			"Content-Type": {"text/csv"},
		},
	}

	err := validateRes(&res)
	if !errors.Is(err, nil) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateResType2(t *testing.T) {
	res := http.Response{
		StatusCode: http.StatusOK,
		Header: map[string][]string{
			"Content-Type": {"text/csv; Charset=utf-8"},
		},
	}

	err := validateRes(&res)
	if !errors.Is(err, nil) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateResInvalidType(t *testing.T) {
	res := http.Response{
		StatusCode: http.StatusOK,
		Header: map[string][]string{
			"Content-Type": {"text/html"},
		},
	}

	err := validateRes(&res)
	if !errors.Is(err, ErrInvalidContentType) {
		t.Fatalf("unexpected error: %v", err)
	}
}
