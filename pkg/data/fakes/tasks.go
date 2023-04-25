package fakes

import (
	"availability/pkg/data"
	"errors"
)

type Source struct {
	ID  int
	URL string
}

type TaskCollection struct {
	Sources []Source
}

func (x TaskCollection) Query(args ...any) (*data.Scanners, error) {
	limit := data.IntArgAt(args, 0)
	if limit == 0 {
		return nil, errors.New("expected limit")
	}

	if limit > len(x.Sources) {
		limit = len(x.Sources)
	}
	res := make([]data.Scanner, 0, limit)
	for i := 0; i < limit; i++ {
		s := TaskScanner{src: x.Sources[i]}
		res = append(res, data.Scanner(&s))
	}
	scanners := data.Scanners(res)
	return &scanners, nil
}

type TaskScanner struct {
	src Source
}

func (x *TaskScanner) Scan(dest ...any) error {
	assign(dest[0], x.src.ID)  // SiteID
	assign(dest[1], x.src.URL) // Domain
	assign(dest[2], 0)         // ProbeID
	assign(dest[3], nil)       // Err
	return nil
}

func assign(dest any, val any) {
	switch d := dest.(type) {
	case *string:
		if d != nil {
			*d = val.(string)
		}
	case *bool:
		if d != nil {
			*d = val.(int) > 0
		}
	case *int32:
		if d != nil {
			*d = int32(val.(int))
		}
	}
}
