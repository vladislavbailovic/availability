package main

import (
	"availability/pkg/data/model"
	"fmt"
	"log"
	"math"
	"strings"
	"time"
)

type renderer interface {
	Render() string
}

type graphMaker interface {
	Make() renderer
}

type graphMeta struct {
	start      time.Time
	end        time.Time
	resolution time.Duration
}

type responseTimesPlotMaker struct {
	graphMeta
	probes []*model.Probe
}

func (x responseTimesPlotMaker) Make() renderer {
	duration := x.end.Sub(x.start)
	frames := float64(duration.Milliseconds()) / float64(x.resolution.Milliseconds())

	var maxTime int
	var minTime int = math.MaxInt
	for _, r := range x.probes {
		dt := int(r.ResponseTime.AsDuration().Milliseconds())
		if dt < minTime {
			minTime = dt
		}
		if dt > maxTime {
			maxTime = dt
		}
	}
	deltaTime := float64(maxTime-minTime) / float64(x.resolution.Milliseconds())
	log.Println("max", maxTime, "min", minTime, "delta", deltaTime)

	points := make([]segment, 0, len(x.probes))
	for _, probe := range x.probes {
		posTime := probe.Recorded.AsTime().Sub(x.start)
		position := float64(posTime.Milliseconds()) / float64(x.resolution.Milliseconds())

		period := probe.ResponseTime.AsDuration()
		length := float64(period.Milliseconds()) / float64(x.resolution.Milliseconds())

		r := point{
			x:     position / frames,
			y:     length / deltaTime,
			label: fmt.Sprintf("%s: %s", probe.Recorded.AsTime(), period),
		}
		points = append(points, segment(r))
	}
	return &svgPointGraph{svgGraph: svgGraph{segments: points}}
}

type incidentReportGraphMaker struct {
	graphMeta
	reports []*model.IncidentReport
}

func (x incidentReportGraphMaker) Make() renderer {
	duration := x.end.Sub(x.start)
	timeframe := float64(duration.Milliseconds()) / float64(x.resolution.Milliseconds())

	blocks := make([]segment, 0, len(x.reports)+int(timeframe))
	for i := 0; i < int(timeframe); i++ {
		t := time.Duration(i) * x.resolution
		r := block{
			x:     float64(i) / timeframe,
			w:     1.0 / timeframe,
			label: fmt.Sprintf("%s: OK", x.start.Add(t)),
			kind:  segmentOK,
		}
		blocks = append(blocks, segment(r))
	}

	for _, report := range x.reports {
		posTime := report.Started.AsTime().Sub(x.start)
		position := float64(posTime.Milliseconds()) / float64(x.resolution.Milliseconds())

		period := report.Ended.AsTime().Sub(report.Started.AsTime())
		length := float64(period.Milliseconds()) / float64(x.resolution.Milliseconds())

		r := block{
			x:     position / timeframe,
			w:     length / timeframe,
			label: fmt.Sprintf("%s: %s", report.Started.AsTime(), period),
			kind:  segmentError,
		}
		blocks = append(blocks, segment(r))
	}
	return &svgBarGraph{svgGraph: svgGraph{segments: blocks}}
}

type segmentType uint8

const (
	segmentNormal segmentType = iota
	segmentOK
	segmentError
)

func (x segmentType) String() string {
	switch x {
	case segmentNormal:
		return ""
	case segmentOK:
		return "ok"
	case segmentError:
		return "error"
	default:
		panic("unknown segment type")
	}
}

type segment interface {
	GetP1() float64
	GetP2() float64
	GetLabel() string
	GetType() segmentType
}

type block struct {
	x, w  float64
	label string
	kind  segmentType
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

func (x block) GetType() segmentType {
	return x.kind
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

func (x point) GetType() segmentType {
	return segmentNormal
}

type svgGraph struct {
	segments []segment
}

type svgBarGraph struct {
	svgGraph
}

func (x *svgBarGraph) Render() string {
	var b strings.Builder
	style := Stylesheet{}

	width := 1000.0
	height := 50.0
	xmlns := "http://www.w3.org/2000/svg"

	fmt.Fprintf(&b, `<svg version="1.1" width="%d" height="%d" xmlns="%s">`,
		int64(width), int64(height*2), xmlns)
	fmt.Fprintf(&b, `<rect x="0" y="0" width="%d" height="%d" class="%s" />`,
		int64(width), int64(height), StylenameMain)

	for _, r := range x.segments {
		x := r.GetP1() * width
		w := r.GetP2() * width
		if w < 1 {
			w = 1
		}
		fmt.Fprintf(&b, `<g class="%s %s">`, StylenameSegment, r.GetType())
		fmt.Fprintf(&b, `<rect x="%f" y="0" width="%f" height="%d" class="period"/>`,
			x, w, int64(height))
		fmt.Fprintf(&b, `<text x="%f" y="%d" class="label">%s</text>`,
			x, int64(height), r.GetLabel()) // TODO: escape/sanitize
		fmt.Fprintf(&b, `</g>`)
	}
	fmt.Fprintf(&b, `<style type="text/css">%s</style>`, style.Render())
	fmt.Fprintf(&b, "</svg>")

	return b.String()
}

type svgPointGraph struct {
	svgGraph
}

func (x *svgPointGraph) Render() string {
	var b strings.Builder
	style := Stylesheet{}

	width := 800.0
	height := 600.0
	xmlns := "http://www.w3.org/2000/svg"

	fmt.Fprintf(&b, `<svg version="1.1" width="%d" height="%d" xmlns="%s">`,
		int64(width+20.0), int64(height+20.0), xmlns)
	fmt.Fprintf(&b, `<rect x="0" y="0" width="%d" height="%d" class="%s" />`,
		int64(width+20.0), int64(height+20.0), StylenameMain)

	var prevX, prevY float64
	for _, r := range x.segments {
		x := r.GetP1() * width
		if x > width {
			x = width
		}
		y := height - (r.GetP2() * height)
		if y > height {
			y = height
		}
		if y < 0 {
			y = 0
		}
		fmt.Fprintf(&b, `<g class="%s %s">`, StylenameSegment, r.GetType())
		fmt.Fprintf(&b, `<circle cx="%f" cy="%f" r="5" class="period"/>`,
			x, y)
		fmt.Fprintf(&b, `<line x1="%f" y1="%f" x2="%f" y2="%f" stroke="black"/>`,
			prevX, prevY, x, y)
		fmt.Fprintf(&b, `<text x="%f" y="%f" class="label">%s</text>`,
			x, y, r.GetLabel()) // TODO: escape/sanitize
		fmt.Fprintf(&b, `</g>`)
		prevX = x
		prevY = y
	}
	fmt.Fprintf(&b, `<style type="text/css">%s</style>`, style.Render())
	fmt.Fprintf(&b, "</svg>")

	return b.String()
}

type Stylename uint16

const (
	StylenameMain Stylename = iota
	StylenameSegment
)

func (x Stylename) String() string {
	switch x {
	case StylenameMain:
		return "main"
	case StylenameSegment:
		return "segment"
	default:
		panic("unknown stylename")
	}
}

type Stylesheet struct{}

func (x Stylesheet) Render() string {
	var b strings.Builder

	fmt.Fprintf(&b, `.%s { fill: white }`, StylenameMain)
	fmt.Fprintf(&b, `.%s.%s .period { fill: green; stroke: black; }`, StylenameSegment, segmentOK)
	fmt.Fprintf(&b, `.%s.%s .period { fill: #cc0000 }`, StylenameSegment, segmentError)
	fmt.Fprintf(&b, `.%s .label { transform: translate(0, 1em); display: none }`, StylenameSegment)
	fmt.Fprintf(&b, `.%s.%s:hover .period { fill: #ff0000 }`, StylenameSegment, segmentError)
	fmt.Fprintf(&b, `.%s:hover .label { display: block }`, StylenameSegment)

	return b.String()
}
