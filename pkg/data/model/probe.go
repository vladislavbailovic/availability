package model

func (x *Probe) IsDown() bool {
	switch x.Err {
	case HttpErr_HTTPERR_NONE, HttpErr_HTTPERR_OK:
		return false
	default:
		return true
	}
}

func (x *ProbeRef) IsDown() bool {
	switch x.Err {
	case HttpErr_HTTPERR_NONE, HttpErr_HTTPERR_OK:
		return false
	default:
		return true
	}
}
