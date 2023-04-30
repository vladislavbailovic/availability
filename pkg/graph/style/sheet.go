package style

import (
	"fmt"
	"strings"

	"availability/pkg/graph/segment"
)

type Sheet struct{}

func (x Sheet) Render() string {
	var b strings.Builder

	fmt.Fprintf(&b, `.%s { fill: white }`, NameMain)
	fmt.Fprintf(&b, `.%s.%s .period { fill: green; stroke: black; }`, NameSegment, segment.OK)
	fmt.Fprintf(&b, `.%s.%s .period { fill: #cc0000 }`, NameSegment, segment.Error)
	fmt.Fprintf(&b, `.%s .label { transform: translate(0, 1em); display: none }`, NameSegment)
	fmt.Fprintf(&b, `.%s.%s:hover .period { fill: #ff0000 }`, NameSegment, segment.Error)
	fmt.Fprintf(&b, `.%s:hover .label { display: block }`, NameSegment)

	return b.String()
}
