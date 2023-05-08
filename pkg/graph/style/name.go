package style

type Name uint16

const (
	NameMain Name = iota
	NameSegment
	NamePeriod
	NameConnector
)

func (x Name) String() string {
	switch x {
	case NameMain:
		return "main"
	case NameSegment:
		return "segment"
	case NamePeriod:
		return "period"
	case NameConnector:
		return "connector"
	default:
		panic("unknown style name")
	}
}
