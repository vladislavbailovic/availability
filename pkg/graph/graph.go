package graph

import "time"

type Renderer interface {
	Render() string
}

type Maker interface {
	Make() Renderer
}

type Meta struct {
	Start      time.Time
	End        time.Time
	Resolution time.Duration
}
