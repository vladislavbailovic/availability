package collections

import (
	"errors"

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