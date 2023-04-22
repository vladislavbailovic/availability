package collections

import (
	"testing"

	"availability/pkg/data/fakes"
)

func Test_GetActiveTasks(t *testing.T) {
	query := &fakes.TaskCollection{
		Sources: []fakes.Source{
			fakes.Source{ID: 1312, URL: "https://snap42.wpmudev.host"},
			fakes.Source{ID: 161, URL: "http://puppychowfoo.rocks"},
		},
	}
	ts, err := GetActiveTasks(query, 10, 10)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(ts) > 10 {
		t.Error("expected at most 10 tasks")
	}
	if len(ts) == 0 {
		t.Error("expected at least some tasks")
	}
	if len(ts) != 2 {
		t.Error("expected exact number of tasks")
	}

	for idx, tk := range ts {
		if tk.Source.SiteID == 0 {
			t.Errorf("missing source SiteID for %d", idx)
		}
	}
}
