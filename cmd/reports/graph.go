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
	blocks := make([]segment, 0, len(x.reports))

	duration := x.end.Sub(x.start)
	timeframe := float64(duration.Milliseconds()) / float64(x.resolution.Milliseconds())
	for _, report := range x.reports {
		posTime := report.Started.AsTime().Sub(x.start)
		position := float64(posTime.Milliseconds()) / float64(x.resolution.Milliseconds())

		period := report.Ended.AsTime().Sub(report.Started.AsTime())
		length := float64(period.Milliseconds()) / float64(x.resolution.Milliseconds())

		r := block{
			x: position / timeframe,
			w: length / timeframe,
		}
		blocks = append(blocks, segment(r))
	}
	return &svg{blocks: blocks}
}

type segment interface {
	GetP1() float64
	GetP2() float64
}

type block struct {
	x, w float64
}

func (x block) GetP1() float64 {
	return x.x
}

func (x block) GetP2() float64 {
	return x.w
}

type svg struct {
	blocks []segment
}

func (x *svg) Render() string {
	var b strings.Builder

	width := 1000.0
	height := 50.0
	xmlns := "http://www.w3.org/2000/svg"

	fmt.Fprintf(&b, `<svg version="1.1" width="%d" height="%d" xmlns="%s">`,
		int64(width), int64(height), xmlns)
	fmt.Fprintf(&b, `<rect x="0" y="0" width="%d" height="%d" fill="green" />`,
		int64(width), int64(height))

	for _, r := range x.blocks {
		x := int64(r.GetP1() * width)
		w := int64(r.GetP2() * width)
		if w < 1 {
			w = 1
		}
		fmt.Fprintf(&b, `<rect x="%d" y="0" width="%d" height="%d" fill="red" />`,
			x, w, int64(height))
	}
	fmt.Fprintf(&b, "</svg>")

	return b.String()
}
