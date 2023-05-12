package collections

import (
	"testing"

	"availability/pkg/data/fakes"
)

func Test_Activate_ErrorsWithInvalidSiteID(t *testing.T) {
	if err := UpdateSource(new(fakes.SourceActivator), 0); err == nil {
		t.Error("expected error")
	}
	if err := UpdateSource(new(fakes.SourceActivator), -1); err == nil {
		t.Error("expected error")
	}
}

func Test_Activate_HappyPath(t *testing.T) {
	if err := UpdateSource(new(fakes.SourceActivator), 1312); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
