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
	//go:embed queries/incident_reports_for_within.sql
	incidentReportsForWithinQuery string
	//go:embed queries/incident_reports_within.sql
	incidentReportsWithinQuery string
)

type IncidentSelection struct {
	sqlConnector
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
	sqlConnector
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
	sqlConnector
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

type IncidentReportCollector struct {
	sqlConnector
}

func (x *IncidentReportCollector) Query(args ...any) (*data.Scanners, error) {
	limit := 100
	siteID := data.IntArgAt(args, 0)
	if siteID == 0 {
		return nil, errors.New("expected siteID")
	}

	since := data.DurationArgAt(args, 1)
	if since == 0 {
		return nil, errors.New("expected period duration")
	}

	if err := x.Connect(); err != nil {
		return nil, err
	}

	stmt, err := x.conn.Prepare(incidentReportsForWithinQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	results, err := stmt.Query(siteID, since.Seconds(), limit)
	if err != nil {
		return nil, err
	}

	res := make([]data.Scanner, 0, limit)
	for i := 0; i < limit; i++ {
		res = append(res, data.Scanner(incidentReportScanner{r: results}))
	}

	scanners := data.Scanners(res)
	return &scanners, nil
}

type IncidentReportPeriodCollector struct {
	sqlConnector
}

func (x *IncidentReportPeriodCollector) Query(args ...any) (*data.Scanners, error) {
	since := data.TimestampArgAt(args, 0)
	if since.Unix() <= 0 {
		return nil, errors.New("expected period timestamp")
	}

	limit := data.IntArgAt(args, 1)
	if limit == 0 {
		return nil, errors.New("expected limit")
	}

	if err := x.Connect(); err != nil {
		return nil, err
	}

	stmt, err := x.conn.Prepare(incidentReportsWithinQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	results, err := stmt.Query(since, limit)
	if err != nil {
		return nil, err
	}

	res := make([]data.Scanner, 0, limit)
	for i := 0; i < limit; i++ {
		res = append(res, data.Scanner(incidentReportScanner{r: results}))
	}

	scanners := data.Scanners(res)
	return &scanners, nil
}

type incidentReportScanner struct {
	r *sql.Rows
}

func (x incidentReportScanner) Scan(dest ...any) error {
	if x.r.Next() {
		return x.r.Scan(dest...)
	}
	return nil
}
