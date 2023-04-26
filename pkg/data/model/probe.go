package model

import (
	"time"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func NewTimeoutProbe(siteID int) *Probe {
	p := new(Probe)
	p.SiteID = int32(siteID)
	p.Recorded = timestamppb.New(time.Now())
	p.Err = HttpErr_HTTPERR_INTERNAL
	p.Msg = "Timeout"
	return p
}

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
