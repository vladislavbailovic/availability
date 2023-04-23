package collections

import (
	"availability/pkg/data"
	"availability/pkg/data/model"
)

type ProbeSet struct {
	probes []*model.Probe
}

func (x *ProbeSet) Add(p *model.Probe) {
	if p != nil && !x.Has(p) {
		x.add(p)
	}
}

func (x *ProbeSet) Has(p *model.Probe) bool {
	for _, n := range x.probes {
		if n.SiteID == p.SiteID && n.Recorded == p.Recorded {
			return true
		}
	}
	return false
}

func (x *ProbeSet) Persist(query data.Inserter) error {
	if len(x.probes) == 0 {
		return nil
	}
	y := make([]any, 0, len(x.probes))
	for _, p := range x.probes {
		y = append(y, p)
	}
	return query.Insert(y...)
}

func (x *ProbeSet) add(p *model.Probe) {
	x.probes = append(x.probes, p)
}
