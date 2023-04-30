package graph

import (
	"availability/pkg/graph/segment"
	"fmt"
)

type SVG struct {
	Segments      []segment.Section
	Height, Width float64
}

func (x SVG) GetHeader() string {
	xmlns := "http://www.w3.org/2000/svg"
	return fmt.Sprintf(`<svg version="1.1" width="%d" height="%d" xmlns="%s">`,
		int64(x.Width), int64(x.Height), xmlns)
}
