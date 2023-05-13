package env

import (
	"os"
	"testing"
)

func Test_Expect_InvalidName(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("expected panic")
		}
	}()
	Variable(1312).Expect()
}

func Test_Expect_MissingValue(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("expected panic")
		}
	}()
	SiteID.Expect()
}

func Test_Expect_HappyPath(t *testing.T) {
	old := os.Getenv(SiteID.String())
	defer func() {
		os.Setenv(SiteID.String(), old)
	}()

	want := "wat"
	os.Setenv(SiteID.String(), want)

	got := SiteID.Expect()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
