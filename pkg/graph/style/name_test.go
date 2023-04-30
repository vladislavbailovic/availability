package style

import "testing"

func Test_Name_String(t *testing.T) {
	if NameMain.String() != "main" {
		t.Error("expected different name")
	}
	if NameSegment.String() != "segment" {
		t.Error("expected different name")
	}

	defer func() {
		if err := recover(); err == nil {
			t.Error("expected panic")
		}
	}()
	n := Name(161)
	if n.String() != "" {
		t.Error("expected empty string")
	}
}
