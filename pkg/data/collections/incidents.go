package collections

import (
	"errors"
	"log"
	"time"

	"availability/pkg/data"
	"availability/pkg/data/model"
)

func GetSiteIncident(query data.Selector, siteID int) (*model.Incident, error) {
	if res, err := query.Query(siteID); err != nil {
		return nil, err
	} else {
		o := new(model.Incident)
		err := res.Scan(
			&o.SiteID,
			&o.DownProbeID,
			&o.UpProbeID)
		return o, err
	}
}

func CloseOffIncident(query data.Updater, o *model.Incident) error {
	if o == nil {
		return errors.New("expected incident")
	}
	if o.SiteID == 0 || o.DownProbeID == 0 || o.UpProbeID == 0 {
		return errors.New("invalid incident")
	}
	return query.Update(o)
}

func CreateNewIncident(query data.Inserter, o *model.Incident) (data.DataID, error) {
	if o == nil {
		return 0, errors.New("expected incident")
	}
	if o.SiteID == 0 || o.DownProbeID == 0 || o.UpProbeID != 0 {
		return 0, errors.New("invalid incident")
	}
	return query.Insert(o)
}

func GetIncidentReportFor(query data.Selector, siteID int) (*model.IncidentReport, error) {
	if res, err := query.Query(siteID); err != nil {
		return nil, err
	} else {
		var started, ended string
		r := new(model.IncidentReport)
		err := res.Scan(
			&r.SiteID,
			&r.URL,
			&started,
			&r.Err,
			&r.Msg,
			&ended)
		r.Started = data.TimestampFromDatetime(started)
		r.Ended = data.TimestampFromDatetime(ended)
		return r, err
	}
}

func GetIncidentReportsFor(query data.Collector, siteID int, since time.Duration) ([]*model.IncidentReport, error) {
	rs := make([]*model.IncidentReport, 0)
	if res, err := query.Query(siteID, since); err != nil {
		return rs, err
	} else {
		for _, result := range *res {
			var started, ended string
			r := new(model.IncidentReport)
			err := result.Scan(
				&r.SiteID,
				&r.URL,
				&started,
				&r.Err,
				&r.Msg,
				&ended)
			if !r.IsValid() {
				continue
			}
			if err != nil {
				log.Printf("WARNING: scan error: %v", err)
				continue
			}
			r.Started = data.TimestampFromDatetime(started)
			r.Ended = data.TimestampFromDatetime(ended)
			rs = append(rs, r)
		}
		return rs, nil
	}
}
