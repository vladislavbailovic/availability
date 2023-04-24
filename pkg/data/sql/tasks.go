package sql

import (
	"database/sql"
	"errors"

	_ "embed"

	"availability/pkg/data"

	_ "github.com/go-sql-driver/mysql"
)

//go:embed queries/active_tasks.sql
var activeTasksQuery string

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
	limit := data.IntArgAt(args, 0)
	if limit == 0 {
		return nil, errors.New("expected limit")
	}
	pingTimeout := data.IntArgAt(args, 1)
	if pingTimeout == 0 {
		return nil, errors.New("expected ping timeout, in seconds")
	}

	if err := x.Connect(); err != nil {
		return nil, err
	}

	stmt, err := x.conn.Prepare(activeTasksQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	results, err := stmt.Query(pingTimeout, limit)
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
