package main

import (
	"availability/pkg/data/model"
	"fmt"
	"strings"
	"time"
)

type renderer interface {
	Render() string
}

type graphMaker interface {
	Make() renderer
}

type incidentReportGraphMaker struct {
	start      time.Time
	end        time.Time
	resolution time.Duration
	reports    []*model.IncidentReport
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
	return &svg{blocks: blocks}
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

type svg struct {
	blocks []segment
}

func (x *svg) Render() string {
	var b strings.Builder
	style := Stylesheet{}

	width := 1000.0
	height := 50.0
	xmlns := "http://www.w3.org/2000/svg"

	fmt.Fprintf(&b, `<svg version="1.1" width="%d" height="%d" xmlns="%s">`,
		int64(width), int64(height*2), xmlns)
	fmt.Fprintf(&b, `<rect x="0" y="0" width="%d" height="%d" class="%s" />`,
		int64(width), int64(height), StylenameMain)

	for _, r := range x.blocks {
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
	fmt.Fprintf(&b, `.%s.%:hover .period { fill: #ff0000 }`, StylenameSegment, segmentError)
	fmt.Fprintf(&b, `.%s:hover .label { display: block }`, StylenameSegment)

	return b.String()
}
