package graph

import (
	"fmt"
	"io"
	"time"

	"github.com/arsmoriendy/opor/gql-srv/csv"
	"github.com/arsmoriendy/opor/gql-srv/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ports       []*model.Port
	lastChecked time.Time
}

func (r *Resolver) GetPorts() {
	rdr, body, err := csv.FetchCsv()
	if err != nil {
		panic(err)
	}
	defer (*body).Close()

	r.lastChecked = time.Now()

	rdr.Read() // skip header
	for {
		rec, err := rdr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			(*body).Close()
			panic(err)
		}

		rdr.FieldsPerRecord = 12

		recpnum, err := csv.ParsePort(rec[1])
		if err != nil {
			(*body).Close()
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
