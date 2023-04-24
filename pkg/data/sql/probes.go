package sql

import (
	"database/sql"
	"strings"

	_ "embed"

	"availability/pkg/data"
	"availability/pkg/data/model"

	_ "github.com/go-sql-driver/mysql"
)

//go:embed queries/insert_probes.sql
var insertProbesQueryPartial string

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
