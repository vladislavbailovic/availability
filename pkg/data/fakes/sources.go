package fakes

import (
	"errors"
)

type SourceActivator struct {
	SiteID int
}

func (x *SourceActivator) Update(v any) error {
	siteID, ok := v.(int)
	if !ok || siteID <= 0 {
		return errors.New("invalid site ID")
	}

	x.SiteID = siteID

	return nil
}
