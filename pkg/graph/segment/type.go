package segment

type Type uint8

const (
	Normal Type = iota
	OK
	Error
)

func (x Type) String() string {
	switch x {
	case Normal:
		return ""
	case OK:
		return "ok"
	case Error:
		return "error"
	default:
		panic("unknown segment type")
	}
}
