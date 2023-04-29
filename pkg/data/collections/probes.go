package collections

import (
	"availability/pkg/data"
	"availability/pkg/data/model"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
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

func (x *ProbeSet) Persist(query data.MultiInserter) (data.DataID, error) {
	if len(x.probes) == 0 {
		return 0, nil
	}
	y := make([]any, 0, len(x.probes))
	for _, p := range x.probes {
		y = append(y, p)
	}
	return query.Insert(y...)
}

func (x *ProbeSet) IsDown() bool {
	for _, p := range x.probes {
		if !p.IsDown() {
			return false
		}
	}
	return true
}

func (x *ProbeSet) add(p *model.Probe) {
	x.probes = append(x.probes, p)
}

func GetProbesForWithin(query data.Collector, siteID int, since time.Duration) ([]*model.Probe, error) {
	ps := make([]*model.Probe, 0)
	if res, err := query.Query(siteID, since); err != nil {
		return ps, err
	} else {
		for _, result := range *res {
			var recorded string
			var responseTime int
			r := new(model.Probe)
			err := result.Scan(
				&r.SiteID,
				&recorded,
				&responseTime,
				&r.Err,
				&r.Msg)
			if err != nil {
				log.Printf("WARNING: scan error: %v", err)
				continue
			}
			r.Recorded = data.TimestampFromDatetime(recorded)
			r.ResponseTime = durationpb.New(time.Duration(responseTime) * time.Millisecond)
			ps = append(ps, r)
		}
		return ps, nil
	}
}
