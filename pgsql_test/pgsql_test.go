package main

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4"
)

const (
	connStr = "postgres://bang:bang@localhost/audit_log?sslmode=disable"
)

func TestTransactionSuccessfulCommit(t *testing.T) {
	t.Parallel()

	conn, err := mustConnectString(connStr)
	if err != nil {
		t.Fatalf("connect pgsql error: %+v", err)
	}
	defer closeConn(conn)

	createSql := `
    create temporary table foo(
      id integer,
      unique (id) initially deferred
    );
  `

	if _, err := conn.Exec(context.Background(), createSql); err != nil {
		t.Fatalf("Failed to create table: %v", err)
	}

	tx, err := conn.Begin(context.Background())
	if err != nil {
		t.Fatalf("conn.Begin failed: %v", err)
	}

	_, err = tx.Exec(context.Background(), "insert into foo(id) values (1)")
	if err != nil {
		t.Fatalf("tx.Exec failed: %v", err)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		t.Fatalf("tx.Commit failed: %v", err)
	}

	var n int64
	err = conn.QueryRow(context.Background(), "select count(*) from foo").Scan(&n)
	if err != nil {
		t.Fatalf("QueryRow Scan failed: %v", err)
	}
	if n != 1 {
		t.Fatalf("Did not receive correct number of rows: %v", n)
	}
}

func TestTxSendBatch(t *testing.T) {
	t.Parallel()

	conn, err := mustConnectString(connStr)
	if err != nil {
		t.Fatalf("connect pgsql error: %+v", err)
	}
	defer closeConn(conn)

	sql := `create temporary table ledger1(
	  id serial primary key,
	  description varchar not null
	);`
	mustExec(conn, sql)

	sql = `create temporary table ledger2(
	  id int primary key,
	  amount int not null
	);`
	mustExec(conn, sql)

	tx, _ := conn.Begin(context.Background())
	batch := &pgx.Batch{}
	batch.Queue("insert into ledger1(description) values($1) returning id", "q1")

	br := tx.SendBatch(context.Background(), batch)

	var id int
	err = br.QueryRow().Scan(&id)
	if err != nil {
		t.Error(err)
	}
	br.Close()

	batch = &pgx.Batch{}
	batch.Queue("insert into ledger2(id,amount) values($1, $2)", id, 2)
	batch.Queue("select amount from ledger2 where id = $1", id)

	br = tx.SendBatch(context.Background(), batch)

	ct, err := br.Exec()
	if err != nil {
		t.Error(err)
	}
	if ct.RowsAffected() != 1 {
		t.Errorf("ct.RowsAffected() => %v, want %v", ct.RowsAffected(), 1)
	}

	var amount int
	err = br.QueryRow().Scan(&amount)
	if err != nil {
		t.Error(err)
	}

	br.Close()
	tx.Commit(context.Background())

	var count int
	conn.QueryRow(context.Background(), "select count(1) from ledger1 where id = $1", id).Scan(&count)
	if count != 1 {
		t.Errorf("count => %v, want %v", count, 1)
	}

	err = br.Close()
	if err != nil {
		t.Fatal(err)
	}

	err = ensureConnValid(conn)
	if err != nil {
		t.Error(err)
	}
}
