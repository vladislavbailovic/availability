package segment

import "testing"

func Test_SegmentType_String(t *testing.T) {
	if Normal.String() != "" {
		t.Error("expected different type")
	}
	if OK.String() != "ok" {
		t.Error("expected different type")
	}
	if Error.String() != "error" {
		t.Error("expected different type")
	}

	defer func() {
		if err := recover(); err == nil {
			t.Error("expected panic")
		}
	}()
	n := Type(161)
	if n.String() != "" {
		t.Error("expected empty string")
	}
}
