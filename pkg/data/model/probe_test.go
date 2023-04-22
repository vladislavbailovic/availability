package model

import "testing"

func Test_IsDown(t *testing.T) {
	p := new(Probe)
	if p.IsDown() {
		t.Error("should be up by default")
	}

	p = &Probe{Err: HttpErr_HTTPERR_INTERNAL}
	if !p.IsDown() {
		t.Error("should be down if error")
	}
}
