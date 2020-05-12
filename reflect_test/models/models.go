package models

import "fmt"

type GorpTx struct {
	sqltx *SqlTx
}

type SqlTx struct {
	txi Tx
}

type Tx interface {
	Commit()
	Rollback()
}

type DriverTx struct {
	GTID string
}

func (t *DriverTx) Commit() {
	fmt.Println("Commit")
}

func (t *DriverTx) Rollback() {
	fmt.Println("Rollback")
}

func (t *DriverTx) GetGTID() string {
	return t.GTID
}

func NewGorpTx(GTID string) *GorpTx {
	return &GorpTx{
		&SqlTx{
			&DriverTx{GTID: GTID},
		},
	}
}

func (gTx *GorpTx) Close() {
	gTx.sqltx.txi = nil
}

type GTIDI interface {
	GetGTID() string
}
