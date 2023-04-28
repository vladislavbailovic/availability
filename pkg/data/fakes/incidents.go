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

type IncidentReportScanner struct {
	Report Report
}

func (x *IncidentReportScanner) Query(args ...any) (data.Scanner, error) {
	res := reportScanner{r: x.Report}
	return data.Scanner(&res), nil
}

type IncidentReportCollector struct {
	Reports []Report
}

func (x *IncidentReportCollector) Query(args ...any) (*data.Scanners, error) {
	siteID := data.IntArgAt(args, 0)
	if siteID == 0 {
		return nil, errors.New("expected siteID")
	}

	res := make([]data.Scanner, 0, len(x.Reports))
	for _, r := range x.Reports {
		s := reportScanner{r: r}
		res = append(res, data.Scanner(&s))
	}
	scanners := data.Scanners(res)
	return &scanners, nil
}

type reportScanner struct {
	r Report
}

func (x *reportScanner) Scan(dest ...any) error {
	assign(dest[0], x.r.ID)      // SiteID
	assign(dest[1], x.r.URL)     // sources.URL
	assign(dest[2], x.r.Started) // Started ts = probe.Recorded
	assign(dest[3], x.r.Err)     // probe.Err
	assign(dest[4], x.r.Msg)     // probe.Msg
	assign(dest[5], x.r.Ended)   // Ended ts = probe.Recorded
	return nil
}

type Report struct {
	ID      int
	URL     string
	Started string
	Msg     string
	Ended   string
	Err     model.HttpErr
}
