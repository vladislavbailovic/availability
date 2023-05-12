package sql

import (
	"database/sql"
	"errors"
	"strings"

	_ "embed"

	"availability/pkg/data"
	"availability/pkg/data/model"

	_ "github.com/go-sql-driver/mysql"
)

var (
	//go:embed queries/insert_probes.sql
	insertProbesQueryPartial string
	//go:embed queries/probes_for_within.sql
	probesForWithinQuery string
)

type ProbeInserter struct {
	sqlConnector
	Probes []*model.Probe
}

func (x *ProbeInserter) Insert(items ...any) (data.DataID, error) {
	if err := x.Connect(); err != nil {
		return 0, err
	}
	vals := make([]string, 0, len(items))
	rpls := make([]any, 0, len(items)*5)
	for _, item := range items {
		if p, ok := item.(*model.Probe); ok {
			vals = append(vals, "(?, ?, ?, ?, ?)")

			rpls = append(rpls, p.SiteID)
			rpls = append(rpls, data.DatetimeToTimestamp(p.Recorded))
			rpls = append(rpls, p.ResponseTime.AsDuration().Milliseconds())
			rpls = append(rpls, p.Err)
			rpls = append(rpls, p.Msg)
		}
	}
	if len(vals) == 0 {
		return 0, nil
	}
	query := insertProbesQueryPartial + strings.Join(vals, ",")
	stmt, err := x.conn.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(rpls...)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return data.DataID(id), err
}

type ProbeCollector struct {
	sqlConnector
}

func (x *ProbeCollector) Query(args ...any) (*data.Scanners, error) {
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

	stmt, err := x.conn.Prepare(probesForWithinQuery)
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
		res = append(res, data.Scanner(probeScanner{r: results}))
	}

	scanners := data.Scanners(res)
	return &scanners, nil
}

type probeScanner struct {
	r *sql.Rows
}

func (x probeScanner) Scan(dest ...any) error {
	if x.r.Next() {
		return x.r.Scan(dest...)
	}
	return nil
}
