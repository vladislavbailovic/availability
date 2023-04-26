package collections

import (
	"strings"
	"testing"

	"availability/pkg/data/fakes"
	"availability/pkg/data/model"
)

func Test_CloseOffIncident_ExpectsValidIncident(t *testing.T) {
	q := new(fakes.IncidentUpdater)
	if err := CloseOffIncident(q, nil); err == nil {
		t.Error("expected incident")
	} else if !strings.Contains(err.Error(), "expected incident") {
		t.Errorf("unexpected error: %v", err)
	}

	p := new(model.Incident)
	if err := CloseOffIncident(q, p); err == nil {
		t.Error("expected incident")
	} else if !strings.Contains(err.Error(), "invalid incident") {
		t.Errorf("unexpected error: %v", err)
	}

	p.SiteID = 161
	if err := CloseOffIncident(q, p); err == nil {
		t.Error("expected incident")
	} else if !strings.Contains(err.Error(), "invalid incident") {
		t.Errorf("unexpected error: %v", err)
	}

	p.DownProbeID = 13
	if err := CloseOffIncident(q, p); err == nil {
		t.Error("expected incident")
	} else if !strings.Contains(err.Error(), "invalid incident") {
		t.Errorf("unexpected error: %v", err)
	}

	p.UpProbeID = 12
	if err := CloseOffIncident(q, p); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_CloseOffIncident_HappyPath(t *testing.T) {
	q := new(fakes.IncidentUpdater)

	p := new(model.Incident)
	p.SiteID = 161
	p.DownProbeID = 13
	p.UpProbeID = 12
	if err := CloseOffIncident(q, p); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if q.Incident.SiteID != p.SiteID ||
		q.Incident.DownProbeID != p.DownProbeID ||
		q.Incident.UpProbeID != p.UpProbeID {
		t.Error("updating went wrong")
	}
}

func Test_CreateNewIncident_ExpectsValidIncident(t *testing.T) {
	q := new(fakes.IncidentInserter)
	if _, err := CreateNewIncident(q, nil); err == nil {
		t.Error("expected incident")
	} else if !strings.Contains(err.Error(), "expected incident") {
		t.Errorf("unexpected error: %v", err)
	}

	p := new(model.Incident)
	if _, err := CreateNewIncident(q, p); err == nil {
		t.Error("expected incident")
	} else if !strings.Contains(err.Error(), "invalid incident") {
		t.Errorf("unexpected error: %v", err)
	}

	p.SiteID = 161
	if _, err := CreateNewIncident(q, p); err == nil {
		t.Error("expected incident")
	} else if !strings.Contains(err.Error(), "invalid incident") {
		t.Errorf("unexpected error: %v", err)
	}

	p.UpProbeID = 13
	if _, err := CreateNewIncident(q, p); err == nil {
		t.Error("expected incident")
	} else if !strings.Contains(err.Error(), "invalid incident") {
		t.Errorf("unexpected error: %v", err)
	}

	p.UpProbeID = 0
	p.DownProbeID = 12
	if _, err := CreateNewIncident(q, p); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_CreateNewIncident_HappyPath(t *testing.T) {
	q := new(fakes.IncidentInserter)

	p := new(model.Incident)
	p.SiteID = 161
	p.DownProbeID = 13

	id, err := CreateNewIncident(q, p)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if q.Incident.SiteID != p.SiteID ||
		q.Incident.DownProbeID != p.DownProbeID ||
		q.Incident.UpProbeID != p.UpProbeID ||
		int32(id) != q.Incident.IncidentID {
		t.Error("updating went wrong")
	}
}
