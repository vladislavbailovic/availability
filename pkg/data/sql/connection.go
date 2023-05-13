package sql

import (
	"database/sql"

	"availability/pkg/env"
)

type sqlConnector struct {
	conn *sql.DB
}

func (x *sqlConnector) Connect() error {
	if x.conn != nil {
		return nil
	}
	db, err := sql.Open("mysql", env.DBConnURI.Expect())
	if err != nil {
		return err
	}
	x.conn = db
	return nil
}

func (x *sqlConnector) Disconnect() {
	if x.conn == nil {
		return
	}
	x.conn.Close()
}
