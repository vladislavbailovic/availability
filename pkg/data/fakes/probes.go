package fakes

import (
	"log"

	"availability/pkg/data/model"
)

type ProbeInserter struct {
	Probes []*model.Probe
}

func (x *ProbeInserter) Insert(items ...any) error {
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
	return nil
}
