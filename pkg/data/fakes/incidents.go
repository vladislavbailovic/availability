package fakes

import (
	"errors"

	"availability/pkg/data"
	"availability/pkg/data/model"
)

type IncidentUpdater struct {
	Incident *model.Incident
}

func (x *IncidentUpdater) Update(v any) error {
	if o, ok := v.(*model.Incident); !ok {
		return errors.New("expected incident")
	} else {
		x.Incident = o
	}
	return nil
}

type IncidentInserter struct {
	Incident *model.Incident
}

func (x *IncidentInserter) Insert(v any) (data.DataID, error) {
	if o, ok := v.(*model.Incident); !ok {
		return 0, errors.New("expected incident")
	} else {
		o.IncidentID = 1312
		x.Incident = o
	}
	return data.DataID(x.Incident.IncidentID), nil
}
