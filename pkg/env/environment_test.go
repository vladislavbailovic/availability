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
	Expect(Name(1312))
}

func Test_Expect_MissingValue(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("expected panic")
		}
	}()
	Expect(SiteID)
}

func Test_Expect_HappyPath(t *testing.T) {
	old := os.Getenv(SiteID.String())
	defer func() {
		os.Setenv(SiteID.String(), old)
	}()

	want := "wat"
	os.Setenv(SiteID.String(), want)

	got := Expect(SiteID)
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
