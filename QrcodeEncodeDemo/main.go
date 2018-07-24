package main

import (
	"fmt"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	err := qrcode.WriteFile("http://www.baidu.com", qrcode.Medium, 256, "qr.png")
	fmt.Println(err)
}
