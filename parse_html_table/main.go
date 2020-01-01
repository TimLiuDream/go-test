package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"log"
	"os"
	"strconv"
)

const (
	sourcePath = "sample/row.html"
	targetPath = "sample/target.html"

	attrRowSpan = "rowspan"
	attrColSpan = "colspan"
)

var (
	rowIndex = -1
	colIndex = -1
)

type cell struct {
	RowIndex int
	ColIndex int
	Value    string
}

func main() {
	file, err := os.Open(sourcePath)
	if err != nil {
		log.Fatalln(err)
	}

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatalln(err)
	}

	sel := doc.Find("tbody")
	if len(sel.Nodes) > 0 {
		rowIndex = -1
		sel.Each(func(i int, s *goquery.Selection) {
			trSel := s.Find("tr")
			if len(trSel.Nodes) > 0 {
				trSel.Each(func(i int, s1 *goquery.Selection) {
					rowIndex++
					colIndex = -1
					tdSel := s1.Find("td")
					if len(tdSel.Nodes) > 0 {
						tdSel.Each(func(i int, s2 *goquery.Selection) {
							colIndex++
							for _, node := range s2.Nodes {
								rowSpan, colSpan := CalculateCellNodeSpan(node.Attr)
								fmt.Println(rowSpan, colSpan)
								fmt.Println(node.Data)
							}
						})
					}
				})
			}
		})
	}
}

// 计算格子节点的行列合并数
func CalculateCellNodeSpan(attrs []html.Attribute) (rowSpan, colSpan int) {
	for _, attr := range attrs {
		if attr.Key == "colspan" {
			col, err := strconv.Atoi(attr.Val)
			if err != nil {
				log.Fatalln(err)
			}
			if col > 1 {
				colSpan = col
			}
		} else if attr.Key == "rowspan" {
			row, err := strconv.Atoi(attr.Val)
			if err != nil {
				log.Fatalln(err)
			}
			if row > 1 {
				rowSpan = row
			}
		}
	}
	return
}
