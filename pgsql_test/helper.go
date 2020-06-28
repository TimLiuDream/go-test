package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

func mustConnectString(connString string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	return conn, err
}

func closeConn(conn *pgx.Conn) error {
	err := conn.Close(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func mustExec(conn *pgx.Conn, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error) {
	return conn.Exec(context.Background(), sql, arguments...)
}

// Do a simple query to ensure the connection is still usable
func ensureConnValid(conn *pgx.Conn) error {
	var sum, rowCount int32

	rows, err := conn.Query(context.Background(), "select generate_series(1,$1)", 10)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var n int32
		rows.Scan(&n)
		sum += n
		rowCount++
	}

	if rows.Err() != nil {
		return fmt.Errorf("conn.Query failed: %v", err)
	}

	if rowCount != 10 {
		return fmt.Errorf("Select called onDataRow wrong number of times")
	}
	if sum != 55 {
		return fmt.Errorf("Wrong values returned")
	}
	return nil
}
