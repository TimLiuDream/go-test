package main

import "fmt"

func main() {
	str:="lkdjfl000"
	bytes:=[]rune(str)
	for from, to := 0, len(bytes)-1; from < to; from, to = from+1, to-1 {
		bytes[from],bytes[to]=bytes[to],bytes[from]
	}
	str = string(bytes)
	fmt.Print(str)
}
