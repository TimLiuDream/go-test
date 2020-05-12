package test_reflect

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/timliudream/golangtraining/reflect_test/models"
)

func TestReflect(t *testing.T) {
	gorpTx := models.NewGorpTx("123456")

	gorpTxValue := reflect.ValueOf(gorpTx)
	gorpTxElem := gorpTxValue.Elem()
	sqlTxField := gorpTxElem.FieldByName("sqltx")
	sqlTxElem := sqlTxField.Elem()
	txiField := sqlTxElem.FieldByName("txi")
	txiFieldType := txiField.Type()
	txiFieldKind := txiField.Kind()
	fmt.Println(txiFieldType)
	fmt.Println(txiFieldKind)
	if txiFieldType.String() != "models.Tx" {
		t.Errorf("txi field type error")
	}
	if txiFieldKind.String() != "interface" {
		t.Errorf("txi field kind error")
	}

	for i := 0; i < txiField.NumMethod(); i++ {
		t.Log(txiField.Method(i).String())
	}
	for i := 0; i < txiField.Elem().NumMethod(); i++ {
		t.Log(txiField.Elem().Method(i).String())
	}
	if txiField.NumMethod() != 2 {
		t.Errorf("txi field method count error")
	}
	if txiField.Elem().NumMethod() != 3 {
		t.Errorf("txi field elem method count error")
	}

	txiFieldElemP := unsafe.Pointer(txiField.Elem().Pointer())
	// pp := txiFieldElemP
	driverTx := (*models.DriverTx)(txiFieldElemP)
	GTID := driverTx.GetGTID()
	if GTID != "123456" {
		t.Error("GTID is error")
	}

	gorpTx.Close()

	rawDriverTx := (*models.DriverTx)(txiFieldElemP)
	GTID = rawDriverTx.GetGTID()
	t.Log(GTID)
	if GTID != "123456" {
		t.Error("GTID is error")
	}
}
