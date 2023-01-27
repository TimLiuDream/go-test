package models

import (
	"fmt"
)

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

func (tx *DriverTx) Commit() {
	fmt.Println("commit")
}

func (tx *DriverTx) Rollback() {
	fmt.Println("rollback")
}

func (tx *DriverTx) GetGTID() string {
	return tx.GTID
}

func NewGorpTx() *GorpTx {
	return &GorpTx{
		sqltx: &SqlTx{
			txi: &DriverTx{
				GTID: "123455",
			},
		},
	}
}

func (tx *GorpTx) Close() {
	tx.sqltx.txi = nil
}
