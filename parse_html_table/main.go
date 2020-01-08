package main

import (
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/timliudream/golangtraining/parse_html_table/models"
	"golang.org/x/net/html"
)

const (
	sourcePath = "sample/row_col.html"
	targetPath = "sample/target.html"

	attrRowSpan = "rowspan"
	attrColSpan = "colspan"
)

var (
	rowIndex = -1
	colIndex = -1
)

func main() {
	file, err := os.Open(sourcePath)
	if err != nil {
		log.Fatalln(err)
	}

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatalln(err)
	}

	mapRowCells := make(map[int][]*models.Cell)
	cells := make([]*models.Cell, 0)

	tableSel := doc.Find("table")
	if len(tableSel.Nodes) > 0 {
		tableSel.Each(func(i int, s *goquery.Selection) {
			rowIndex = -1
			rowSel := s.Find("tr")
			if len(rowSel.Nodes) > 0 {
				rowSel.Each(func(i int, s *goquery.Selection) {
					rowIndex++
					colIndex = -1
					colSel := s.Find("td")
					if len(colSel.Nodes) > 0 {
						colSel.Each(func(i int, s *goquery.Selection) {
							colIndex++
							for _, node := range s.Nodes {
								rowSpan, colSpan := CalculateCellNodeSpan(node.Attr)
								if rowSpan == 0 && colSpan == 0 { // 没有合并单元格
									c := models.AddCell(mapRowCells, rowIndex, colIndex, node.FirstChild.Data)
									cells = append(cells, c)
									mapRowCells[rowIndex] = append(mapRowCells[rowIndex], c)
								} else if rowSpan > 0 && colSpan == 0 { // 只有行合并
									c := models.AddCell(mapRowCells, rowIndex, colIndex, node.FirstChild.Data)
									cells = append(cells, c)
									mapRowCells[rowIndex] = append(mapRowCells[rowIndex], c)
									for i := 1; i < rowSpan; i++ {
										c := models.AddCell(mapRowCells, rowIndex+i, colIndex, node.FirstChild.Data)
										cells = append(cells, c)
										mapRowCells[rowIndex+i] = append(mapRowCells[rowIndex+i], c)
									}
								} else if rowSpan == 0 && colSpan > 0 { // 只有列合并
									c := models.AddCell(mapRowCells, rowIndex, colIndex, node.FirstChild.Data)
									cells = append(cells, c)
									mapRowCells[rowIndex] = append(mapRowCells[rowIndex], c)
									for i := 1; i < colSpan; i++ {
										c := models.AddCell(mapRowCells, rowIndex, colIndex+i, node.FirstChild.Data)
										cells = append(cells, c)
										mapRowCells[rowIndex] = append(mapRowCells[rowIndex], c)
									}
								} else { // 行列合并
									for i := 0; i < rowSpan; i++ {
										for j := 0; j < colSpan; j++ {
											c := models.AddCell(mapRowCells, rowIndex+i, colIndex+j, node.FirstChild.Data)
											cells = append(cells, c)
											mapRowCells[rowIndex+i] = append(mapRowCells[rowIndex+i], c)
										}
									}
								}
							}
						})
					}
				})
			}
		})
	}
	sort.Sort(models.CellSorter(cells))
	tbodyStr := buildNewTable(mapRowCells, cells)
	tableSel.ReplaceWithHtml(tbodyStr)
	htmlStr, err := doc.Html()
	if err != nil {
		log.Fatalln(err)
	}
	_ = ioutil.WriteFile(targetPath, []byte(htmlStr), 0644)
}

func buildNewTable(mapRowCells map[int][]*models.Cell, cells []*models.Cell) string {
	maxRow := 0
	for _, cell := range cells {
		if cell.RowIndex > maxRow {
			maxRow = cell.RowIndex
		}
	}
	sb := new(strings.Builder)
	sb.WriteString("<table>")
	sb.WriteString("<tbody>")
	for i := 0; i <= maxRow; i++ {
		sb.WriteString("<tr>")
		partCells := mapRowCells[i]
		sort.Sort(models.CellSorter(partCells))
		for _, partCell := range partCells {
			sb.WriteString("<td>")
			sb.WriteString(partCell.Value)
			sb.WriteString("</td>")
		}
		sb.WriteString("</tr>")
	}
	sb.WriteString("</tbody>")
	sb.WriteString("</table>")
	return sb.String()
}

// 计算格子节点的行列合并数
func CalculateCellNodeSpan(attrs []html.Attribute) (rowSpan, colSpan int) {
	for _, attr := range attrs {
		if attr.Key == attrColSpan {
			col, err := strconv.Atoi(attr.Val)
			if err != nil {
				log.Fatalln(err)
			}
			if col > 1 {
				colSpan = col
			}
		} else if attr.Key == attrRowSpan {
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
