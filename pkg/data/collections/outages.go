package collections

import (
	"availability/pkg/data"
	"availability/pkg/data/model"
	"log"
)

func GetLastSiteOutage(query data.Selector, siteID int) (*model.Outage, error) {
	if res, err := query.Query(siteID); err != nil {
		return nil, err
	} else {
		o := new(model.Outage)
		err := res.Scan(
			&o.SiteID,
			&o.DownProbeID,
			&o.UpProbeID)
		log.Printf("Scanned outage: SiteID: %d, DPI: %d, UPI: %d", o.SiteID, o.DownProbeID, o.UpProbeID)
		return o, err
	}
}
