package style

import (
	"fmt"
	"strings"

	"availability/pkg/graph/segment"
)

type Sheet struct{}

func (x Sheet) Render() string {
	var b strings.Builder

	fmt.Fprint(&b, `<style type="text/css" media="all">`)

	fmt.Fprintf(&b, `.%s { fill: white }`, NameMain)
	fmt.Fprintf(&b, `.%s.%s .period { fill: green; stroke: #007700; }`, NameSegment, segment.OK)
	fmt.Fprintf(&b, `.%s.%s .period { fill: #ff0000 }`, NameSegment, segment.Error)

	fmt.Fprintf(&b, `.%s .label { transform: translate(0, 1em); }`, NameSegment)
	fmt.Fprintf(&b, `.%s .label tspan { display: none }`, NameSegment)
	fmt.Fprintf(&b, `.%s .label tspan.short { display: block; fill: #666 }`, NameSegment)

	// If we can hover:
	fmt.Fprint(&b, `@media (hover:hover) {`)
	// Label toggle on hover
	fmt.Fprintf(&b, `.%s:hover .label tspan.long { display: block }`, NameSegment)
	fmt.Fprintf(&b, `.%s:hover .label tspan.short { display: none }`, NameSegment)
	// Error fade-in on hover
	fmt.Fprintf(&b, `.%s.%s .period { fill: #cc0000 }`, NameSegment, segment.Error)
	fmt.Fprintf(&b, `.%s.%s:hover .period { fill: #ff0000 }`, NameSegment, segment.Error)
	fmt.Fprint(&b, `}`)

	fmt.Fprint(&b, `</style>`)

	return b.String()
}
