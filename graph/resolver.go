package graph

import (
	c "encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/arsmoriendy/opor/gql-srv/csv"
	"github.com/arsmoriendy/opor/gql-srv/graph/model"
	"github.com/arsmoriendy/opor/gql-srv/internal"
	"github.com/arsmoriendy/opor/gql-srv/internal/loglvl"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ports       []*model.Port
	lastChecked time.Time
}

func New() (r Resolver) {
	// TODO: refresh periodically
	r.refreshPorts()
	return
}

func (r *Resolver) refreshPorts() {
	var rdr *c.Reader
	var body io.ReadCloser
	var err error

	if internal.IsDevMode() {
		f, err := os.Open("test/data/service-names-port-numbers.csv")
		if errors.Is(err, os.ErrNotExist) {
			panic(fmt.Errorf("%w: make sure to run in project root", err))
		}
		if err != nil {
			panic(err)
		}
		if loglvl.LogLvl >= loglvl.INFO {
			log.Println("using test/mock csv file")
		}
		rdr = c.NewReader(f)
		body = f
	} else {
		if loglvl.LogLvl >= loglvl.INFO {
			log.Println("fetching csv...")
		}
		rdr, body, err = csv.FetchCsv()
	}

	if err != nil {
		panic(err)
	}
	defer body.Close()

	r.lastChecked = time.Now()

	rdr.Read() // skip header
	for {
		rec, err := rdr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			body.Close()
			panic(err)
		}

		rdr.FieldsPerRecord = 12

		recpnum, err := csv.ParsePort(rec[1])
		if err != nil {
			body.Close()
			panic(fmt.Errorf("%w: with values %v", err, rec))
		}
		r.ports = append(r.ports, &model.Port{
			ServiceName:             &rec[0],
			PortNumber:              recpnum,
			TransportProtocol:       &rec[2],
			Description:             &rec[3],
			Assignee:                &rec[4],
			Contact:                 &rec[5],
			RegistrationDate:        &rec[6],
			ModificationDate:        &rec[7],
			Reference:               &rec[8],
			ServiceCode:             &rec[9],
			UnauthorizedUseReported: &rec[10],
			AssignmentNotes:         &rec[11],
		})

	}
}

func (r *Resolver) SearchPort(toFind *model.Port) (bool, uint, error) {
	n := uint(len(r.ports))
	if n == 0 {
		return false, 0, internal.ErrSearchEmptyArr
	}

	return r.bsp(toFind, 0, n-1, n/2)
}

// binary search for array of ports
func (r *Resolver) bsp(toFind *model.Port, startIdx uint, endIdx uint, pivotIdx uint) (bool, uint, error) {
	pivotVal := r.ports[pivotIdx]

	isLarger, err := pivotVal.Larger(toFind)
	if err != nil {
		return false, 0, err
	}
	isSmaller, err := pivotVal.Smaller(toFind)
	if err != nil {
		return false, 0, err
	}
	isEqual, err := pivotVal.Equal(toFind)
	if err != nil {
		return false, 0, err
	}

	if endIdx-startIdx < 2 {
		return isEqual, pivotIdx, nil
	}

	if isLarger {
		return r.bsp(toFind, startIdx, pivotIdx, startIdx+(pivotIdx-startIdx)/2)
	} else if isSmaller {
		pivotIdx++
		return r.bsp(toFind, pivotIdx, endIdx, pivotIdx+(endIdx-pivotIdx)/2)
	} else {
		return true, pivotIdx, nil
	}
}
