package collections

import (
	"availability/pkg/data"
	"availability/pkg/data/model"
	"log"
)

// TODO: implement fetching pings
// This is going to be something like:
// SELECT * FROM sites WHERE toPing=1 AND somehow-last-pinged WITHIN <PING_INTERVAL+1>
func GetActiveTasks(query data.Collector, limit int) ([]*model.Task, error) {
	ts := make([]*model.Task, 0, limit)

	if res, err := query.Query(limit); err != nil {
		return ts, err
	} else if res == nil {
		return ts, nil
	} else {
		for _, r := range *res {
			s := new(model.Source)
			p := new(model.Probe)
			err := r.Scan(
				&s.SiteID,
				&s.URL,
				&p.Err)
			if err != nil {
				log.Printf("WARNING: scan error: %v", err)
				continue
			}
			if s.SiteID == 0 {
				continue
			}
			t := new(model.Task)
			t.Source = s
			ts = append(ts, t)
		}
	}

	return ts, nil
}
