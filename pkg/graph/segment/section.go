package segment

type Section interface {
	GetP1() float64
	GetP2() float64
	GetLabel() string
	GetType() Type
}
