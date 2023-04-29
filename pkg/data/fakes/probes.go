package fakes

import (
	"errors"
	"log"

	"availability/pkg/data"
	"availability/pkg/data/model"
)

type ProbeInserter struct {
	Probes []*model.Probe
}

func (x *ProbeInserter) Insert(items ...any) (data.DataID, error) {
	for _, item := range items {
		if p, ok := item.(*model.Probe); ok {
			log.Printf("Persisting probe for %d:", p.SiteID)
			log.Printf("\t- Recorded: %s", p.Recorded)
			log.Printf("\t- Duration: %s", p.ResponseTime)
			log.Printf("\t- Err: %d", p.Err)
			log.Printf("\t- Msg: %s", p.Msg)
			x.Probes = append(x.Probes, p)
		}
	}
	return data.DataID(len(x.Probes)), nil
}

type Probe struct {
	SiteID, ResponseTime, Err int
	Msg, Recorded             string
}

type ProbeCollector struct {
	Probes []Probe
}

func (x *ProbeCollector) Query(args ...any) (*data.Scanners, error) {
	siteID := data.IntArgAt(args, 0)
	if siteID == 0 {
		return nil, errors.New("expected siteID")
	}

	res := make([]data.Scanner, 0, len(x.Probes))
	for _, r := range x.Probes {
		s := probeScanner{p: r}
		res = append(res, data.Scanner(&s))
	}
	scanners := data.Scanners(res)
	return &scanners, nil
}

type probeScanner struct {
	p Probe
}

func (x *probeScanner) Scan(dest ...any) error {
	assign(dest[0], x.p.SiteID)
	assign(dest[1], x.p.Recorded)
	assign(dest[2], x.p.ResponseTime)
	assign(dest[3], x.p.Err)
	assign(dest[4], x.p.Msg)
	return nil
}
