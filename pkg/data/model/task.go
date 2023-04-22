package model

func (x *Task) WasPreviouslyDown() bool {
	if x.Previous == nil {
		return false
	}
	return x.Previous.IsDown()
}
