package collections

import (
	"errors"

	"availability/pkg/data"
	"availability/pkg/data/model"
)

func GetSiteOutage(query data.Selector, siteID int) (*model.Outage, error) {
	if res, err := query.Query(siteID); err != nil {
		return nil, err
	} else {
		o := new(model.Outage)
		err := res.Scan(
			&o.SiteID,
			&o.DownProbeID,
			&o.UpProbeID)
		return o, err
	}
}

func CloseOffOutage(query data.Updater, o *model.Outage) error {
	if o == nil {
		return errors.New("expected outage")
	}
	if o.SiteID == 0 || o.DownProbeID == 0 || o.UpProbeID == 0 {
		return errors.New("invalid outage")
	}
	return query.Update(o)
}

func CreateNewOutage(query data.Inserter, o *model.Outage) (data.DataID, error) {
	if o == nil {
		return 0, errors.New("expected outage")
	}
	if o.SiteID == 0 || o.DownProbeID == 0 || o.UpProbeID != 0 {
		return 0, errors.New("invalid outage")
	}
	return query.Insert(o)
}
