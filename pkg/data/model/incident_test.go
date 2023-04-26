package model

import "testing"

func Test_NewIncident(t *testing.T) {
	x := NewIncident(1312, 161)

	if x.SiteID != 1312 {
		t.Errorf("expected site id: %d", x.SiteID)
	}
	if x.DownProbeID != 161 {
		t.Errorf("expected down probe id: %d", x.DownProbeID)
	}
	if x.UpProbeID != 0 {
		t.Errorf("unexpected up probe id: %d", x.UpProbeID)
	}
}

func Test_Incident_Close(t *testing.T) {
	x := NewIncident(13, 12)
	if x.UpProbeID != 0 {
		t.Errorf("unexpected up probe id: %d", x.UpProbeID)
	}

	x.Close(161)
	if x.UpProbeID != 161 {
		t.Errorf("expected up probe id: %d", x.UpProbeID)
	}
}
