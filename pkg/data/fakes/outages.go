package fakes

import (
	"errors"

	"availability/pkg/data"
	"availability/pkg/data/model"
)

type OutageUpdater struct {
	Outage *model.Outage
}

func (x *OutageUpdater) Update(v any) error {
	if o, ok := v.(*model.Outage); !ok {
		return errors.New("expected outage")
	} else {
		x.Outage = o
	}
	return nil
}

type OutageInserter struct {
	Outage *model.Outage
}

func (x *OutageInserter) Insert(v any) (data.DataID, error) {
	if o, ok := v.(*model.Outage); !ok {
		return 0, errors.New("expected outage")
	} else {
		o.OutageID = 1312
		x.Outage = o
	}
	return data.DataID(x.Outage.OutageID), nil
}
