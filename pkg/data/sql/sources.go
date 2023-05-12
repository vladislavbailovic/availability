package sql

import (
	"errors"

	"availability/pkg/data/model"

	_ "embed"
)

var (
	//go:embed queries/update_source_query.sql
	updateSourceQuery string
)

type SourceActivator struct {
	sqlConnector
}

func (x *SourceActivator) Update(v any) error {
	siteID, ok := v.(int)
	if !ok || siteID <= 0 {
		return errors.New("invalid site ID")
	}

	stmt, err := x.conn.Prepare(updateSourceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(model.SourceActive, siteID)
	return err
}

type SourceDeactivator struct {
	sqlConnector
}

func (x *SourceDeactivator) Update(v any) error {
	siteID, ok := v.(int)
	if !ok || siteID <= 0 {
		return errors.New("invalid site ID")
	}

	stmt, err := x.conn.Prepare(updateSourceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(model.SourceInactive, siteID)
	return err
}
