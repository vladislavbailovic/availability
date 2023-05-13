package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_ExtractNumberFromPathAt(t *testing.T) {
	suite := map[string]int{
		"/nopenopenope":   0,
		"/activate/1312":  1312,
		"/activate/1312/": 1312,
		"/whatever/dude":  0,
	}
	for test, want := range suite {
		t.Run(test, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, test, nil)
			got := ExtractNumberFromPathAt(req, 1)
			if want != got {
				t.Errorf("want %d, got %d", want, got)
			}
		})
	}
}
