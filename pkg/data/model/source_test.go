package model

import "testing"

func Test_Source_IsValid(t *testing.T) {
	valid := &Source{SiteID: 1312, URL: "test"}
	if !valid.IsValid() {
		t.Error("expected valid")
	}

	invalid := new(Source)
	if invalid.IsValid() {
		t.Error("expected invalid")
	}
}

func Test_NewSource_IsValid(t *testing.T) {
	valid := &NewSource{SiteID: 1312, URL: "test"}
	if !valid.IsValid() {
		t.Error("expected valid")
	}

	invalid := new(NewSource)
	if invalid.IsValid() {
		t.Error("expected invalid")
	}
}
