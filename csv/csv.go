// Helper for parsing IANA csv
package csv

import (
	"encoding/csv"
	"errors"
	"io"
	"net/http"
	"strconv"
)

var ErrInvalidStatusCode = errors.New("csv: invalid response status code (not 200 OK)")

func FetchCsv() (*csv.Reader, *io.ReadCloser, error) {
	res, err := http.Get("https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.csv")
	if err != nil {
		return nil, nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, nil, ErrInvalidStatusCode
	}

	return csv.NewReader(res.Body), &res.Body, nil
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
