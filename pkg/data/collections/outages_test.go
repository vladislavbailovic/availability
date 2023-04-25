package collections

import (
	"strings"
	"testing"

	"availability/pkg/data/fakes"
	"availability/pkg/data/model"
)

func Test_CloseOffOutage_ExpectsValidOutage(t *testing.T) {
	q := new(fakes.OutageUpdater)
	if err := CloseOffOutage(q, nil); err == nil {
		t.Error("expected outage")
	} else if !strings.Contains(err.Error(), "expected outage") {
		t.Errorf("unexpected error: %v", err)
	}

	p := new(model.Outage)
	if err := CloseOffOutage(q, p); err == nil {
		t.Error("expected outage")
	} else if !strings.Contains(err.Error(), "invalid outage") {
		t.Errorf("unexpected error: %v", err)
	}

	p.SiteID = 161
	if err := CloseOffOutage(q, p); err == nil {
		t.Error("expected outage")
	} else if !strings.Contains(err.Error(), "invalid outage") {
		t.Errorf("unexpected error: %v", err)
	}

	p.DownProbeID = 13
	if err := CloseOffOutage(q, p); err == nil {
		t.Error("expected outage")
	} else if !strings.Contains(err.Error(), "invalid outage") {
		t.Errorf("unexpected error: %v", err)
	}

	p.UpProbeID = 12
	if err := CloseOffOutage(q, p); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_CloseOffOutage_HappyPath(t *testing.T) {
	q := new(fakes.OutageUpdater)

	p := new(model.Outage)
	p.SiteID = 161
	p.DownProbeID = 13
	p.UpProbeID = 12
	if err := CloseOffOutage(q, p); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if q.Outage.SiteID != p.SiteID ||
		q.Outage.DownProbeID != p.DownProbeID ||
		q.Outage.UpProbeID != p.UpProbeID {
		t.Error("updating went wrong")
	}
}

func Test_CreateNewOutage_ExpectsValidOutage(t *testing.T) {
	q := new(fakes.OutageInserter)
	if _, err := CreateNewOutage(q, nil); err == nil {
		t.Error("expected outage")
	} else if !strings.Contains(err.Error(), "expected outage") {
		t.Errorf("unexpected error: %v", err)
	}

	p := new(model.Outage)
	if _, err := CreateNewOutage(q, p); err == nil {
		t.Error("expected outage")
	} else if !strings.Contains(err.Error(), "invalid outage") {
		t.Errorf("unexpected error: %v", err)
	}

	p.SiteID = 161
	if _, err := CreateNewOutage(q, p); err == nil {
		t.Error("expected outage")
	} else if !strings.Contains(err.Error(), "invalid outage") {
		t.Errorf("unexpected error: %v", err)
	}

	p.UpProbeID = 13
	if _, err := CreateNewOutage(q, p); err == nil {
		t.Error("expected outage")
	} else if !strings.Contains(err.Error(), "invalid outage") {
		t.Errorf("unexpected error: %v", err)
	}

	p.UpProbeID = 0
	p.DownProbeID = 12
	if _, err := CreateNewOutage(q, p); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_CreateNewOutage_HappyPath(t *testing.T) {
	q := new(fakes.OutageInserter)

	p := new(model.Outage)
	p.SiteID = 161
	p.DownProbeID = 13

	id, err := CreateNewOutage(q, p)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if q.Outage.SiteID != p.SiteID ||
		q.Outage.DownProbeID != p.DownProbeID ||
		q.Outage.UpProbeID != p.UpProbeID ||
		int32(id) != q.Outage.OutageID {
		t.Error("updating went wrong")
	}
}
