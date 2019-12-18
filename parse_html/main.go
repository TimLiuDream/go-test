package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	data, err := ioutil.ReadFile("example.html")
	if err != nil {
		log.Fatalln(err)
	}
	buf := bytes.NewBufferString(string(data))
	doc, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		log.Fatalln(err)
	}
	_ = doc.Find("div[class=ones-toc-info]").Remove()
	result, err := doc.Html()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)
}
