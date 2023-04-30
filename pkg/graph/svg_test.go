package graph

import (
	"strings"
	"testing"
)

func Test_SVG_GetHeader(t *testing.T) {
	x := SVG{Height: 161.0, Width: 1312.0}
	r := x.GetHeader()

	if !strings.Contains(r, `width="1312"`) {
		t.Errorf("expected width: %q", r)
	}

	if !strings.Contains(r, `height="161"`) {
		t.Errorf("expected height: %q", r)
	}
}
