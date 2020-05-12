package main

import (
	"fmt"
)

const (
	x = iota
	y
	z
	w
)

const v = iota

const (
	h, i, j = iota, iota, iota
)

const (
	a       = iota
	b       = "B"
	c       = iota
	d, e, f = iota, iota, iota
	g       = iota
)

type clientFlag uint32

const (
	clientLongPassword clientFlag = 1 << iota
	clientFoundRows
	clientLongFlag
	clientConnectWithDB
	clientNoSchema
	clientCompress
	clientODBC
	clientLocalFiles
	clientIgnoreSpace
	clientProtocol41
	clientInteractive
	clientSSL
	clientIgnoreSIGPIPE
	clientTransactions
	clientReserved
	clientSecureConn
	clientMultiStatements
	clientMultiResults
	clientPSMultiResults
	clientPluginAuth
	clientConnectAttrs
	clientPluginAuthLenEncClientData
	clientCanHandleExpiredPasswords
	clientSessionTrack
	clientDeprecateEOF
)

func main() {
	func2()
}

func func1() {
	fmt.Print(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v)
}

func func2() {
	fmt.Println(clientLongPassword | clientFoundRows | clientLongFlag | clientConnectWithDB | clientNoSchema | clientCompress | clientODBC | clientLocalFiles | clientIgnoreSpace | clientProtocol41 | clientInteractive | clientIgnoreSIGPIPE | clientTransactions | clientReserved | clientSecureConn)
}
