package style

type Name uint16

const (
	NameMain Name = iota
	NameSegment
)

func (x Name) String() string {
	switch x {
	case NameMain:
		return "main"
	case NameSegment:
		return "segment"
	default:
		panic("unknown style name")
	}
}
