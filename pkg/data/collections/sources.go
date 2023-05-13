package collections

import (
	"errors"
	"net/url"

	"availability/pkg/data"
	"availability/pkg/data/model"
)

func UpdateSource(query data.Updater, siteID data.DataID) error {
	if !siteID.IsValid() {
		return errors.New("invalid site ID")
	}
	return query.Update(siteID.ToNumericID())
}

func CreateNewSource(query data.Inserter, src *model.NewSource) (data.DataID, error) {
	if !src.IsValid() {
		return 0, errors.New("expected valid source data")
	}
	if _, err := url.ParseRequestURI(src.URL); err != nil {
		return 0, err
	}

	return query.Insert(src)
}
