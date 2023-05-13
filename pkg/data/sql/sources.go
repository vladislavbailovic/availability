package sql

import (
	"errors"

	"availability/pkg/data"
	"availability/pkg/data/model"

	_ "embed"
)

var (
	//go:embed queries/update_source_query.sql
	updateSourceQuery string

	//go:embed queries/insert_source_query.sql
	insertSourceQuery string
)

type SourceActivator struct {
	sqlConnector
}

func (x *SourceActivator) Update(v any) error {
	siteID, ok := v.(int)
	if !ok || siteID <= 0 {
		return errors.New("invalid site ID")
	}

	if err := x.Connect(); err != nil {
		return err
	}
	defer x.Disconnect()

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

	if err := x.Connect(); err != nil {
		return err
	}
	defer x.Disconnect()

	stmt, err := x.conn.Prepare(updateSourceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(model.SourceInactive, siteID)
	return err
}

type SourceInserter struct {
	sqlConnector
}

func (x *SourceInserter) Insert(v any) (data.DataID, error) {
	site, ok := v.(*model.NewSource)
	if !ok {
		return 0, errors.New("expected new source")
	}

	if err := x.Connect(); err != nil {
		return 0, err
	}
	defer x.Disconnect()

	stmt, err := x.conn.Prepare(insertSourceQuery)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(site.SiteID, site.URL)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return data.DataID(id), err
}
