package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/timliudream/go-test/parse_html_table2/constants"
	"github.com/timliudream/go-test/parse_html_table2/models"
	"github.com/timliudream/go-test/parse_html_table2/utils"
)

func main() {
	doc, err := readSourceHtml("sample/row.html")
	if err != nil {
		log.Fatalln(err)
	}

	// Traversal html table element
	doc.Find("table").Each(func(i int, tableSelection *goquery.Selection) {
		// Traversal html table body element
		tableSelection.Find("tbody").Each(func(i int, tbodySelection *goquery.Selection) {
			mapRowColCell := make(map[int]map[int]*models.Cell)
			rowIndex := -1
			// Traversal html table body row element
			tbodySelection.Find("tr").Each(func(i int, trSelection *goquery.Selection) {
				rowIndex++
				colIndex := -1
				// trHtml := "<tr>"
				// Traversal html table body header and cell element
				trSelection.Children().Each(func(i int, cellSelection *goquery.Selection) {
					colIndex++
					node := cellSelection.Nodes[0]

					// 判断当前格子在mapRowCells能否找到
					mapColCell, foundMapColCell := mapRowColCell[rowIndex]
					if foundMapColCell {
						cell, foundCell := mapColCell[colIndex]
						if foundCell {
							cellHtml := utils.BuildCellHtml(cell.MapAttr, cell.Value)
							// trHtml += cellHtml
							cellSelection.BeforeHtml(cellHtml)
							// ret, _ := doc.Html()
							// fmt.Println(ret)

							// cellSelection.SetHtml(cellHtml)
							// colIndex++
						}
					}

					mapAttr := utils.MapAttr(node.Attr)
					rowSpan := 0
					colSpan := 0
					rawRowSpan, rowSpanFound := mapAttr[constants.AttrRowSpan]
					if rowSpanFound {
						rowSpan, err = strconv.Atoi(rawRowSpan)
						if err != nil {
							log.Fatalln(err)
						}
					}
					rawColSpan, colSpanFound := mapAttr[constants.AttrColSpan]
					if colSpanFound {
						colSpan, err = strconv.Atoi(rawColSpan)
						if err != nil {
							log.Fatalln(err)
						}
					}
					if rowSpan == 0 && colSpan == 0 { // 没有合并单元格
						cellHtml := utils.BuildCellHtml(mapAttr, node.FirstChild.Data)
						cellSelection.SetHtml(cellHtml)
						// trHtml += cellHtml
					} else if rowSpan > 0 && colSpan == 0 { // 只有行合并
						cellSelection.RemoveAttr(constants.AttrRowSpan)
						value := node.FirstChild.Data
						cellHtml := utils.BuildCellHtml(mapAttr, value)
						cellSelection.SetHtml(cellHtml)
						// trHtml += cellHtml
						for i := 1; i < rowSpan; i++ {
							c := models.NewCell(rowIndex+i, colIndex, node.FirstChild.Data, mapAttr)
							_, found := mapRowColCell[rowIndex+i]
							if !found {
								mapColCell := make(map[int]*models.Cell)
								mapColCell[colIndex] = c
								mapRowColCell[rowIndex+i] = mapColCell
							} else {
								mapRowColCell[rowIndex+i][colIndex] = c
							}
						}
					} else if rowSpan == 0 && colSpan > 0 { // 只有列合并

					} else { // 行列合并

					}
				})

				// trHtml += "</tr>"
				// fmt.Println(trHtml)
				// trSelection.ReplaceWithHtml(trHtml)
				// ret, _ := trSelection.Html()
				// fmt.Println(ret)
			})
		})
	})

	err = saveTargetHtml(doc, "sample/target.html")
	if err != nil {
		log.Fatalln(err)
	}
}

// readSourceHtml read source html into doc
func readSourceHtml(sourcePath string) (doc *goquery.Document, err error) {
	file, err := os.Open(sourcePath)
	if err != nil {
		log.Fatalln(err)
	}

	return goquery.NewDocumentFromReader(file)
}

// saveTargetHtml save new target html
func saveTargetHtml(doc *goquery.Document, targetPath string) error {
	htmlStr, err := doc.Html()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(targetPath, []byte(htmlStr), 0644)
}
