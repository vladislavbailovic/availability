package sql

import (
	"database/sql"
	"strings"

	"availability/pkg/data"
	"availability/pkg/data/model"

	_ "github.com/go-sql-driver/mysql"
)

type ProbeInserter struct {
	conn   *sql.DB
	Probes []*model.Probe
}

func (x *ProbeInserter) Connect() error {
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

func (x *ProbeInserter) Disconnect() {
	if x.conn == nil {
		return
	}
	x.conn.Close()
}

func (x *ProbeInserter) Insert(items ...any) error {
	if err := x.Connect(); err != nil {
		return err
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
		return nil
	}
	query := "INSERT INTO probes (site_id, recorded, response_time, err, msg) VALUES" +
		strings.Join(vals, ",")
	stmt, err := x.conn.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(rpls...)
	if err != nil {
		return err
	}
	return nil
}
