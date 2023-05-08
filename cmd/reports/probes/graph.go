package probes

import (
	"fmt"
	"html/template"
	"math"
	"strings"
	"time"

	"availability/pkg/data/model"
	"availability/pkg/graph"
	"availability/pkg/graph/segment"
	"availability/pkg/graph/style"
)

const curveDelta float64 = 10

type rawPoint struct {
	t  time.Time
	ds []int64
}

type ResponseTimesPlot struct {
	graph.Meta
	Probes []*model.Probe
}

func (x ResponseTimesPlot) Make() graph.Renderer {
	return x.makeProbesWithinResolutionPlot()
}

// Plots resolution averages
// makeProbesWithinResolutionPlot generates a renderer for the response times
// plot
//
// This function generates a graph of the response time data, grouped by time
// interval specified in the resolution field of x.  It loops through the
// probes, and for each probe, checks if it falls in the time interval. If it
// does, it adds the response time of the probe to the corresponding rawPoint.
// RawPoints represent a unique interval determined by the resolution
// property.  The final graph is generated by iterating through all the
// rawPoints, calculating the mean response time and generating a point for
// that interval on the graph.
//
// Returns a renderer interface, which is used by the exporter to write the
// graph to the output file.
func (x ResponseTimesPlot) makeProbesWithinResolutionPlot() graph.Renderer {
	res := x.Resolution.Milliseconds()
	frames := x.End.Sub(x.Start).Milliseconds() / res
	points := make([]segment.Section, 0, frames)

	// data is a map containing a rawPoint for each interval
	data := map[int64]*rawPoint{}

	// maxTime and minTime stores the highest and lowest response times respectively
	var maxTime int64
	var minTime int64 = math.MaxInt64

	// assert x.start > x.end - res
	// assert res > probe interval
	for i := x.Start.UnixMilli(); i < x.End.UnixMilli()-res; i += res {
		for _, p := range x.Probes {
			dt := p.ResponseTime.AsDuration().Milliseconds()
			if dt < minTime {
				minTime = dt
			}
			if dt > maxTime {
				maxTime = dt
			}
			// check if probe falls in the current interval
			t := p.Recorded.AsTime()
			if t.UnixMilli() > i && t.UnixMilli() <= i+res {
				// if there's already a rawPoint for this interval, add the response time to the slice
				// else, create a new rawPoint for this interval and add the response time to its slice
				if pt, ok := data[i]; ok {
					pt.ds = append(pt.ds, dt)
				} else {
					r := new(rawPoint)
					r.t = t
					r.ds = []int64{dt}
					data[i] = r
				}
			}
		}
	}

	deltaTime := float64(maxTime-minTime) / float64(res)

	// iterate through all the rawPoints, calculate the mean response time and
	// generate a point for that interval on the graph
	for i := x.Start.UnixMilli(); i < x.End.UnixMilli()-res; i += res {
		if rp, ok := data[i]; ok {
			posTime := rp.t.Sub(x.Start)
			position := float64(posTime.Milliseconds()) / float64(res)
			var sum int64
			for _, d := range rp.ds {
				sum += d
			}
			rts := sum / int64(len(rp.ds))
			length := float64(rts) / float64(res)
			r := point{
				x:     position / float64(frames),
				y:     length / deltaTime,
				label: fmt.Sprintf("%s: %dms over %d probes", rp.t, rts, len(rp.ds)),
			}
			points = append(points, segment.Section(r))
		}
	}
	return &svgPointGraph{SVG: graph.SVG{Segments: points, Height: 600.0, Width: 800.0}}
}

// Plots all points individually
// makeAllProbesPlot generates an SVG point graph for the response times of a
// set of probes over a given time range using the specified resolution. It
// calculates the frames, maximum and minimum response times, and the time
// difference between each frame. It then creates points for each probe, where
// the x-position represents the time frame and y-position represents the
// response time, normalized by the time difference between frames. A label is
// also generated for each point, displaying the recorded time and response
// period.
func (x ResponseTimesPlot) makeAllProbesPlot() graph.Renderer {
	res := x.Resolution.Milliseconds()
	duration := x.End.Sub(x.Start)
	frames := float64(duration.Milliseconds()) / float64(res)

	var maxTime int
	var minTime int = math.MaxInt
	for _, r := range x.Probes {
		dt := int(r.ResponseTime.AsDuration().Milliseconds())
		if dt < minTime {
			minTime = dt
		}
		if dt > maxTime {
			maxTime = dt
		}
	}
	deltaTime := float64(maxTime-minTime) / float64(res)

	points := make([]segment.Section, 0, len(x.Probes))
	for _, probe := range x.Probes {
		posTime := probe.Recorded.AsTime().Sub(x.Start)
		position := float64(posTime.Milliseconds()) / float64(res)

		period := probe.ResponseTime.AsDuration()
		length := float64(period.Milliseconds()) / float64(res)

		r := point{
			x:     position / frames,
			y:     length / deltaTime,
			label: fmt.Sprintf("%s: %s", probe.Recorded.AsTime(), period),
		}
		points = append(points, segment.Section(r))
	}
	return &svgPointGraph{SVG: graph.SVG{Segments: points, Height: 600.0, Width: 800.0}}
}

type point struct {
	x, y  float64
	label string
}

func (x point) GetP1() float64 {
	return x.x
}

func (x point) GetP2() float64 {
	return x.y
}

func (x point) GetLabel() string {
	return x.label
}

func (x point) GetType() segment.Type {
	return segment.Normal
}

type svgPointGraph struct {
	graph.SVG
}

func (g *svgPointGraph) Render() string {
	var b strings.Builder
	sheet := style.Sheet{}

	fmt.Fprintf(&b, g.GetHeader())
	fmt.Fprintf(&b, `<rect x="0" y="0" width="%d" height="%d" class="%s" />`,
		int64(g.Width), int64(g.Height), style.NameMain)

	radius := 10
	dfr := math.Pi + 1.0 // to center the short label; TODO: why though?

	var prevX, prevY float64
	var initX, initY float64
	var path, pts strings.Builder
	for idx, r := range g.Segments {
		x := r.GetP1() * g.Width
		if x > g.Width {
			x = g.Width
		}
		y := g.Height - (r.GetP2() * g.Height)
		if y > g.Height {
			y = g.Height
		}
		if y < 0 {
			y = 0
		}

		if initX == 0.0 {
			initX = x
			prevX = x
		}
		if initY == 0.0 {
			initY = y
			prevY = y
		}

		fmt.Fprintf(&pts, `<g class="%s %s">`, style.NameSegment, r.GetType())
		fmt.Fprintf(&pts, `<circle cx="%f" cy="%f" r="%d" class="%s"/>`,
			x, y, radius, style.NamePeriod)
		fmt.Fprintf(&pts, `<text x="%f" y="%f" class="label">`, x, y)
		fmt.Fprintf(&pts, `<tspan x="%f" y="%f" class="short">%d</tspan>`, x-dfr, y+dfr, idx+1)
		fmt.Fprintf(&pts, `<tspan class="long">%s</tspan>`, template.HTMLEscapeString(r.GetLabel()))
		fmt.Fprint(&pts, `</text>`)
		fmt.Fprint(&pts, `</g>`)

		// TODO: improve curve smoothing
		if math.Abs(x-prevX) > curveDelta && math.Abs(y-prevY) > curveDelta {
			fmt.Fprintf(&path, "C %f,%f %f,%f %f,%f\n", x, prevY, prevX, y, x, y)
		} else {
			fmt.Fprintf(&path, "L %f,%f\n", x, y)
		}

		prevX = x
		prevY = y
	}
	fmt.Fprintf(&b, `<path d="M %f,%f %s" fill="none" class="%s" />`,
		initX, initY, path.String(), style.NameConnector)
	fmt.Fprintf(&b, pts.String())
	b.WriteString(sheet.Render())
	fmt.Fprintf(&b, "</svg>")

	return b.String()
}
