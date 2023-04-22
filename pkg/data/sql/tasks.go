package sql

import (
	"database/sql"
	"errors"

	"availability/pkg/data"

	_ "github.com/go-sql-driver/mysql"
)

type TaskCollection struct {
	conn *sql.DB
}

func (x *TaskCollection) Connect() error {
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

func (x *TaskCollection) Disconnect() {
	if x.conn == nil {
		return
	}
	x.conn.Close()
}

func (x *TaskCollection) Query(args ...any) (*data.Scanners, error) {
	var limit int
	if len(args) > 0 {
		if l, ok := args[0].(int); ok {
			limit = l
		}
	}
	if limit == 0 {
		return nil, errors.New("expected limit")
	}

	if err := x.Connect(); err != nil {
		return nil, err
	}

	stmt, err := x.conn.Prepare(
		"SELECT sources.site_id, url, IFNULL(err, 0) as err FROM sources LEFT JOIN (" +
			"SELECT site_id, err, recorded FROM probes ORDER BY recorded DESC LIMIT 1" +
			") AS probe ON sources.site_id=probe.site_id " +
			"WHERE sources.active=1 AND " +
			"probe.recorded BETWEEN TIMESTAMPADD(MINUTE, -2, NOW()) AND NOW() " +
			"LIMIT ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	results, err := stmt.Query(limit)
	if err != nil {
		return nil, err
	}

	res := make([]data.Scanner, 0, limit)
	for i := 0; i < limit; i++ {
		res = append(res, data.Scanner(rowScanner{r: results}))
	}

	scanners := data.Scanners(res)
	return &scanners, nil
}

type rowScanner struct {
	r *sql.Rows
}

func (x rowScanner) Scan(dest ...any) error {
	if x.r.Next() {
		return x.r.Scan(dest...)
	}
	return nil
}
