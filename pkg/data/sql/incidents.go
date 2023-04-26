package sql

import (
	"database/sql"
	"errors"
	"log"

	"availability/pkg/data"
	"availability/pkg/data/model"

	_ "embed"

	_ "github.com/go-sql-driver/mysql"
)

var (
	//go:embed queries/last_site_incident.sql
	lastSiteIncidentQuery string
	//go:embed queries/update_incident.sql
	updateIncidentQuery string
	//go:embed queries/insert_incident.sql
	insertIncidentQuery string
)

type IncidentSelection struct {
	conn *sql.DB
}

func (x *IncidentSelection) Connect() error {
	if x.conn != nil {
		return nil
	}
	db, err := sql.Open("mysql", "root:root@tcp(avbl-data:3306)/narfs")
	if err != nil {
		return err
	}
	x.conn = db
	return nil
}

func (x *IncidentSelection) Disconnect() {
	if x.conn == nil {
		return
	}
	x.conn.Close()
}

func (x *IncidentSelection) Query(args ...any) (data.Scanner, error) {
	siteID := data.IntArgAt(args, 0)
	if siteID == 0 {
		return nil, errors.New("expected site ID")
	}

	if err := x.Connect(); err != nil {
		return nil, err
	}

	stmt, err := x.conn.Prepare(lastSiteIncidentQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(siteID)
	return incidentSelectionScanner{row}, nil
}

type IncidentUpdater struct {
	conn *sql.DB
}

func (x *IncidentUpdater) Connect() error {
	if x.conn != nil {
		return nil
	}
	db, err := sql.Open("mysql", "root:root@tcp(avbl-data:3306)/narfs")
	if err != nil {
		return err
	}
	x.conn = db
	return nil
}

func (x *IncidentUpdater) Disconnect() {
	if x.conn == nil {
		return
	}
	x.conn.Close()
}

func (x *IncidentUpdater) Update(v any) error {
	o, ok := v.(*model.Incident)
	if !ok {
		return errors.New("expected incident")
	}

	if err := x.Connect(); err != nil {
		return err
	}

	stmt, err := x.conn.Prepare(updateIncidentQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(o.UpProbeID, o.SiteID, o.DownProbeID)
	return err
}

type IncidentInserter struct {
	conn *sql.DB
}

func (x *IncidentInserter) Connect() error {
	if x.conn != nil {
		return nil
	}
	db, err := sql.Open("mysql", "root:root@tcp(avbl-data:3306)/narfs")
	if err != nil {
		return err
	}
	x.conn = db
	return nil
}

func (x *IncidentInserter) Disconnect() {
	if x.conn == nil {
		return
	}
	x.conn.Close()
}

func (x *IncidentInserter) Insert(v any) (data.DataID, error) {
	o, ok := v.(*model.Incident)
	if !ok {
		return 0, errors.New("expected incident")
	}

	if err := x.Connect(); err != nil {
		return 0, err
	}

	stmt, err := x.conn.Prepare(insertIncidentQuery)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(o.SiteID, o.DownProbeID)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return data.DataID(id), err
}

type incidentSelectionScanner struct {
	r *sql.Row
}

func (x incidentSelectionScanner) Scan(dest ...any) error {
	err := x.r.Scan(dest...)
	if err != nil {
		log.Println(err)
		if err != sql.ErrNoRows {
			return nil
		}
		return err
	}
	return nil
}
