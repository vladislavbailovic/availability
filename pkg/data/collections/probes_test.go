package collections

import (
	"testing"
	"time"

	"availability/pkg/data/fakes"
	"availability/pkg/data/model"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_ProbeSet_Basic(t *testing.T) {
	x := new(ProbeSet)
	x.Add(nil)
	if len(x.probes) != 0 {
		t.Error("should not accept nil")
	}

	p := new(model.Probe)
	p.SiteID = 1312
	x.Add(p)
	if len(x.probes) != 1 {
		t.Error("should add valid probe")
	}

	x.Add(p)
	if len(x.probes) != 1 {
		t.Error("should not add probe twice")
	}

	p1 := new(model.Probe)
	p1.SiteID = 1312
	p1.Recorded = timestamppb.New(time.Now())
	x.Add(p1)
	if len(x.probes) != 2 {
		t.Errorf("should add new probe: %v", x.probes)
	}
}

func Test_ProbeSet_Persist(t *testing.T) {
	x := new(ProbeSet)

	p := new(model.Probe)
	p.SiteID = 1312
	x.Add(p)

	p1 := new(model.Probe)
	p1.SiteID = 1312
	p1.Recorded = timestamppb.New(time.Now())
	x.Add(p1)

	query := new(fakes.ProbeInserter)
	if err := x.Persist(query); err != nil {
		t.Error("expected successful insert")
	}

	if len(x.probes) != len(query.Probes) {
		t.Errorf("want %d, got %d", len(x.probes), len(query.Probes))
	}
}
