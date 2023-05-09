package incidents

import (
	"fmt"
	"strings"
	"text/template"
	"time"

	"availability/pkg/data/model"
	"availability/pkg/graph"
	"availability/pkg/graph/segment"
	"availability/pkg/graph/style"
)

type ReportGraph struct {
	graph.Meta
	Reports []*model.IncidentReport
}

func (x ReportGraph) Make() graph.Renderer {
	duration := x.End.Sub(x.Start)
	timeframe := float64(duration.Milliseconds()) / float64(x.Resolution.Milliseconds())

	blocks := make([]segment.Section, 0, len(x.Reports)+int(timeframe))
	for i := 0; i < int(timeframe); i++ {
		t := time.Duration(i) * x.Resolution
		r := block{
			x:     float64(i) / timeframe,
			w:     1.0 / timeframe,
			label: fmt.Sprintf("%s: OK", x.Start.Add(t)),
			kind:  segment.OK,
		}
		blocks = append(blocks, segment.Section(r))
	}

	for _, report := range x.Reports {
		posTime := report.Started.AsTime().Sub(x.Start)
		position := float64(posTime.Milliseconds()) / float64(duration.Milliseconds())

		period := report.Ended.AsTime().Sub(report.Started.AsTime())
		length := float64(period.Milliseconds()) / float64(duration.Milliseconds())

		r := block{
			x:     position,
			w:     length,
			label: fmt.Sprintf("%s: %s", report.Started.AsTime(), period),
			kind:  segment.Error,
		}
		blocks = append(blocks, segment.Section(r))
	}
	return &svgBarGraph{SVG: graph.SVG{Segments: blocks, Height: 75.0, Width: 1000.0}}
}

type block struct {
	x, w  float64
	label string
	kind  segment.Type
}

func (x block) GetP1() float64 {
	return x.x
}

func (x block) GetP2() float64 {
	return x.w
}

func (x block) GetLabel() string {
	return x.label
}

func (x block) GetType() segment.Type {
	return x.kind
}

type svgBarGraph struct {
	graph.SVG
}

func (g *svgBarGraph) Render() string {
	var b strings.Builder
	sheet := style.Sheet{}

	fmt.Fprintf(&b, g.GetHeader())
	fmt.Fprintf(&b, `<rect x="0" y="0" width="%d" height="%d" class="%s" />`,
		int64(g.Width), int64(g.Height), style.NameMain)

	for idx, r := range g.Segments {
		x := r.GetP1() * g.Width
		if x < 0 {
			x = 0
		}
		w := r.GetP2() * g.Width
		if w < 1 {
			w = 1
		}
		if w > g.Width {
			w = g.Width
		}
		fmt.Fprintf(&b, `<g class="%s %s">`, style.NameSegment, r.GetType())
		fmt.Fprintf(&b, `<rect x="%f" y="0" width="%f" height="%d" class="period"/>`,
			x, w, int64(g.Height))
		fmt.Fprintf(&b, `<text x="%f" y="%d" class="label">`, x, int64(g.Height))
		fmt.Fprintf(&b, `<tspan x="%f" y="%f" class="short">%d</tspan>`, x, (g.Height), idx+1)
		fmt.Fprintf(&b, `<tspan x="0" y="%f" class="long">%s</tspan>`, g.Height, template.HTMLEscapeString(r.GetLabel()))
		fmt.Fprint(&b, `</text>`)
		fmt.Fprint(&b, `</g>`)
	}
	b.WriteString(sheet.Render())
	fmt.Fprintf(&b, "</svg>")

	return b.String()
}
