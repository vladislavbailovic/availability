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
	//go:embed queries/last_site_outage.sql
	lastSiteOutageQuery string
	//go:embed queries/update_outage.sql
	updateOutageQuery string
	//go:embed queries/insert_outage.sql
	insertOutageQuery string
)

type OutageSelection struct {
	conn *sql.DB
}

func (x *OutageSelection) Connect() error {
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

func (x *OutageSelection) Disconnect() {
	if x.conn == nil {
		return
	}
	x.conn.Close()
}

func (x *OutageSelection) Query(args ...any) (data.Scanner, error) {
	siteID := data.IntArgAt(args, 0)
	if siteID == 0 {
		return nil, errors.New("expected site ID")
	}

	if err := x.Connect(); err != nil {
		return nil, err
	}

	stmt, err := x.conn.Prepare(lastSiteOutageQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(siteID)
	return outageSelectionScanner{row}, nil
}

type OutageUpdater struct {
	conn *sql.DB
}

func (x *OutageUpdater) Connect() error {
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

func (x *OutageUpdater) Disconnect() {
	if x.conn == nil {
		return
	}
	x.conn.Close()
}

func (x *OutageUpdater) Update(v any) error {
	o, ok := v.(*model.Outage)
	if !ok {
		return errors.New("expected outage")
	}

	if err := x.Connect(); err != nil {
		return err
	}

	stmt, err := x.conn.Prepare(updateOutageQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(o.UpProbeID, o.SiteID, o.DownProbeID)
	return err
}

type OutageInserter struct {
	conn *sql.DB
}

func (x *OutageInserter) Connect() error {
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

func (x *OutageInserter) Disconnect() {
	if x.conn == nil {
		return
	}
	x.conn.Close()
}

func (x *OutageInserter) Insert(v any) (data.DataID, error) {
	o, ok := v.(*model.Outage)
	if !ok {
		return 0, errors.New("expected outage")
	}

	if err := x.Connect(); err != nil {
		return 0, err
	}

	stmt, err := x.conn.Prepare(insertOutageQuery)
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

type outageSelectionScanner struct {
	r *sql.Row
}

func (x outageSelectionScanner) Scan(dest ...any) error {
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
