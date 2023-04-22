package collections

import "testing"

func Test_GetActiveTasks(t *testing.T) {
	f := new(FakeTaskCollection)
	ts, err := GetActiveTasks(f, 10)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(ts) > 10 {
		t.Error("expected at most 10 tasks")
	}
	if len(ts) == 0 {
		t.Error("expected at least some tasks")
	}

	for idx, tk := range ts {
		if tk.Source.SiteID == 0 {
			t.Errorf("missing source SiteID for %d", idx)
		}
	}
}
