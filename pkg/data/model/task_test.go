package model

import "testing"

func Test_WasPreviouslyDown(t *testing.T) {
	task := new(Task)
	if task.WasPreviouslyDown() {
		t.Error("should previously be up by default")
	}

	p := ProbeRef{Err: HttpErr_HTTPERR_INTERNAL}
	task.Previous = &p
	if !task.WasPreviouslyDown() {
		t.Error("expected previously to be down when previous probe err set")
	}
}
