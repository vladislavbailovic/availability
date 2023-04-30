package style

import (
	"availability/pkg/graph/segment"
	"strings"
	"testing"
)

func Test_Sheet_Render(t *testing.T) {
	s := Sheet{}
	r := s.Render()
	suite := []string{
		NameMain.String(),
		NameSegment.String(),
		segment.Error.String(),
	}
	for _, test := range suite {
		t.Run(test, func(t *testing.T) {
			if !strings.Contains(r, test) {
				t.Errorf("want %q but could not find it in %q", test, r)
			}
		})
	}
}
