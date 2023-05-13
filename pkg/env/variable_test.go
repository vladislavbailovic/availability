package env

import "testing"

func Test_NameString(t *testing.T) {
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
