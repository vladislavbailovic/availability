package sql

import (
	"database/sql"
	"os"

	"availability/pkg/env"
)

type sqlConnector struct {
	conn *sql.DB
}

func (x *sqlConnector) Connect() error {
	if x.conn != nil {
		return nil
	}
	db, err := sql.Open("mysql", os.Getenv(env.DBConnURI.String()))
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
