package model

import (
	"errors"

	"github.com/arsmoriendy/opor/gql-srv/internal"
)

var ErrEmptyPortNumber = errors.New("empty PortNumber")

// Checks if p's PortNumber contains a portNumber.
// Returns wheather or not portNumber is found and the index of it.
func (p *Port) Contains(portNumber int) (bool, uint, error) {
	return internal.BinarySearch(&p.PortNumber, portNumber)
}

func (p *Port) Empty() bool {
	return len(p.PortNumber) == 0
}

func (s *Port) Larger(p *Port) (bool, error) {
	if s.Empty() || p.Empty() {
		return false, ErrEmptyPortNumber
	}
	sfirst := s.PortNumber[0]
	plast := p.PortNumber[len(p.PortNumber)-1]
	return sfirst > plast, nil
}

func (s *Port) Smaller(p *Port) (bool, error) {
	if s.Empty() || p.Empty() {
		return false, ErrEmptyPortNumber
	}
	slast := s.PortNumber[len(s.PortNumber)-1]
	plast := p.PortNumber[len(p.PortNumber)-1]
	return slast > plast, nil
}

func (s *Port) Equal(p *Port) (bool, error) {
	sempty := s.Empty()
	pempty := p.Empty()
	if sempty && pempty {
		return true, nil
	}
	if sempty || pempty {
		return false, ErrEmptyPortNumber
	}
	for _, ppn := range p.PortNumber {
		found, _, _ := internal.BinarySearch(&s.PortNumber, ppn)
		if found {
			return true, nil
		}
	}
	return false, nil
}
