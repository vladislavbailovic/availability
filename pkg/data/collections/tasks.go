package collections

import (
	"availability/pkg/data"
	"availability/pkg/data/model"
	"errors"
	"fmt"
	"math/rand"
)

// TODO: implement fetching pings
// This is going to be something like:
// SELECT * FROM sites WHERE toPing=1 AND somehow-last-pinged WITHIN <PING_INTERVAL+1>
func GetActiveTasks(query data.Collector, limit int) ([]*model.Task, error) {
	ts := make([]*model.Task, 0, limit)

	if res, err := query.Query(limit); err != nil {
		return ts, err
	} else if res == nil {
		return ts, nil
	} else {
		for _, r := range *res {
			s := new(model.Source)
			p := new(model.Probe)
			err := r.Scan(
				&s.SiteID,
				&s.URL,
				&p.Err)
			if err != nil {
				continue
			}
			t := new(model.Task)
			t.Source = s
			ts = append(ts, t)
		}
	}

	return ts, nil
}

type FakeTaskCollection struct{}

func (x FakeTaskCollection) Query(args ...any) (*data.Scanners, error) {
	var limit int
	if len(args) > 0 {
		if l, ok := args[0].(int); ok {
			limit = l
		}
	} else {
		limit = 10
	}
	if limit == 0 {
		return nil, errors.New("expected limit")
	}
	res := make([]data.Scanner, 0, limit)
	for i := 0; i < limit; i++ {
		res = append(res, data.Scanner(new(FakeTaskScanner)))
	}
	scanners := data.Scanners(res)
	return &scanners, nil
}

type FakeTaskScanner struct{}

func (x *FakeTaskScanner) Scan(dest ...any) error {
	siteId := rand.Intn(1312)
	assign(dest[0], siteId)                                    // SiteID
	assign(dest[1], fmt.Sprintf("http://site-%d.com", siteId)) // Domain
	assign(dest[2], nil)                                       // Err
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
