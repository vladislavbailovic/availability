package collections

import (
	"errors"

	"availability/pkg/data"
)

func UpdateSource(query data.Updater, siteID int) error {
	if siteID <= 0 {
		return errors.New("invalid site ID")
	}
	return query.Update(siteID)
}
