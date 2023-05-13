package env

import (
	"os"
	"testing"
)

func Test_VariableString(t *testing.T) {
	suite := map[string]Variable{
		"AVBL_SITE_ID":         SiteID,
		"AVBL_SITE_URL":        SiteURL,
		"AVBL_PREVIOUSLY_DOWN": PreviouslyDown,
		"AVBL_DBCONN_URI":      DBConnURI,
	}
	for want, test := range suite {
		t.Run(want, func(t *testing.T) {
			got := test.String()
			if want != got {
				t.Errorf("want %q, got %q", want, got)
			}
		})
	}
}

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

func Test_WithFallback_GetsFallback(t *testing.T) {
	if "fallback" != SiteID.WithFallback("fallback") {
		t.Errorf("unexpected fallback value: %q",
			SiteID.WithFallback("fallback"))
	}
}

func Test_WithFallback_HappyPath(t *testing.T) {
	old := os.Getenv(SiteID.String())
	defer func() {
		os.Setenv(SiteID.String(), old)
	}()

	want := "wat"
	os.Setenv(SiteID.String(), want)

	got := SiteID.WithFallback("fallback")
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
