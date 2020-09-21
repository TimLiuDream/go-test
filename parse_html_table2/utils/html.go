package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/timliudream/go-test/parse_html_table2/constants"
	"golang.org/x/net/html"
)

// MapAttr map html node attributes
func MapAttr(attrs []html.Attribute) (mapAttr map[string]string) {
	mapAttr = make(map[string]string)
	for _, attr := range attrs {
		mapAttr[attr.Key] = attr.Val
	}
	return mapAttr
}

// BuildCellHtml build cell html string to replace
func BuildCellHtml(mapAttr map[string]string, value string) (cellHtml string) {
	sb := new(strings.Builder)
	sb.WriteString("<td")
	for key, val := range mapAttr {
		// need remove rowspan and colspan attributes
		if key == constants.AttrRowSpan || key == constants.AttrColSpan {
			continue
		}
		sb.WriteString(" ")
		sb.WriteString(key)
		sb.WriteString("=")
		sb.WriteString(fmt.Sprintf("\"%s\"", val))
	}
	sb.WriteString(">")
	sb.WriteString(value)
	sb.WriteString("</td>")
	return sb.String()
}

// CalculateCellNodeSpan calculate cell span
func CalculateCellNodeSpan(attrs []html.Attribute) (rowSpan, colSpan int) {
	for _, attr := range attrs {
		if attr.Key == constants.AttrColSpan {
			col, err := strconv.Atoi(attr.Val)
			if err != nil {
				log.Fatalln(err)
			}
			if col > 1 {
				colSpan = col
			}
		} else if attr.Key == constants.AttrRowSpan {
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
