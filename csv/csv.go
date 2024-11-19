// Helper for parsing IANA csv
package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var ErrInvalidStatusCode = errors.New("csv: invalid response status code (not 200 OK)")
var ErrInvalidContentType = errors.New("csv: invalid response content type")

func FetchCsv() (*csv.Reader, io.ReadCloser, error) {
	res, err := http.Get("https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.csv")
	if err != nil {
		return nil, nil, err
	}

	// check status code
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, nil, ErrInvalidStatusCode
	}

	// check content type
	contentType := res.Header.Get("Content-Type")
	contentTypeSlices := strings.Split(contentType, ";")
	found := false
	for _, s := range contentTypeSlices {
		trimmed := strings.Trim(s, " ")
		if trimmed == "text/csv" {
			found = true
			break
		}
	}
	if !found {
		return nil, nil, fmt.Errorf(
			"%w: found '%s', expected 'text/csv'",
			ErrInvalidContentType,
			contentType,
		)
	}

	return csv.NewReader(res.Body), res.Body, nil
}

// Wrapper for determining if a port field is:
// - empty
// - a range
// - a singular number
func ParsePort(port string) ([]int, error) {
	// for entries with empty port field?
	if port == "" {
		return []int{}, nil
	}

	rport, err := strconv.Atoi(port)
	rports := []int{rport}
	if err != nil {
		rports, err = parsePortRange(port)
		if err != nil {
			return nil, err
		}
	}
	return rports, nil
}

func parsePortRange(portRange string) ([]int, error) {
	startStr, endStr, parseEnd := "", "", false

	for _, c := range portRange {
		if c == '-' {
			parseEnd = true
			continue
		}
		if parseEnd {
			endStr = endStr + string(c)
			continue
		}
		startStr = startStr + string(c)
	}

	start, err := strconv.Atoi(startStr)
	if err != nil {
		return nil, err
	}
	end, err := strconv.Atoi(endStr)
	if err != nil {
		return nil, err
	}

	return makeRange(start, end), nil
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
