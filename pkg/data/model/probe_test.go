package model

import "testing"

func Test_NewTimeoutProbe(t *testing.T) {
	p := NewTimeoutProbe(1312)

	if p.SiteID != 1312 {
		t.Errorf("unexpected site ID: %v", p.SiteID)
	}
	if p.Recorded == nil {
		t.Error("expected recorded time")
	}
	if p.Msg == "" {
		t.Error("expected an error message")
	}
	if p.Err != HttpErr_HTTPERR_INTERNAL {
		t.Errorf("expected internal error: %v", p.Err)
	}
}

func Test_Probe_IsDown(t *testing.T) {
	p := new(Probe)
	if p.IsDown() {
		t.Error("should be up by default")
	}

	p = &Probe{Err: HttpErr_HTTPERR_INTERNAL}
	if !p.IsDown() {
		t.Error("should be down if error")
	}
}

func Test_ProbeRef_IsDown(t *testing.T) {
	p := new(ProbeRef)
	if p.IsDown() {
		t.Error("should be up by default")
	}

	p = &ProbeRef{Err: HttpErr_HTTPERR_INTERNAL}
	if !p.IsDown() {
		t.Error("should be down if error")
	}
}
