package main

import (
	"fmt"
	"log"
	"math"

	"github.com/tealeg/xlsx"
)

func main() {
	path := "/Users/tim/Documents/test.xlsx"

	xlFile, err := xlsx.OpenFile(path)
	if err != nil {
		log.Fatal(err)
	}
	mergeCellValue := ""
	hMergeScope := make([]int, 2) //行合并的范围
	hMergeScope[0] = math.MaxInt8
	vMergeScope := make([]int, 2)
	vMergeScope[0] = math.MaxInt8

	for _, sheet := range xlFile.Sheets {
		for rowIndex, row := range sheet.Rows {
			if rowIndex < vMergeScope[0] || rowIndex > vMergeScope[1] {
				mergeCellValue = ""
				hMergeScope[0] = math.MaxInt8
			}
			cells := row.Cells
			for cellIndex := 0; cellIndex < len(cells); cellIndex++ {
				if cellIndex >= hMergeScope[0] && cellIndex <= hMergeScope[1] {
					fmt.Printf("行：%d,列：%d，值:%s\n", rowIndex, cellIndex, mergeCellValue)
				} else {
					text := cells[cellIndex].String()

					h := cells[cellIndex].HMerge
					if h > 0 {
						mergeCellValue = text
						maxHMerge := cellIndex + h
						hMergeScope[0] = cellIndex
						hMergeScope[1] = maxHMerge
					}
					v := cells[cellIndex].VMerge
					if v > 0 {
						mergeCellValue = text
						maxVMerge := rowIndex + v
						vMergeScope[0] = rowIndex
						vMergeScope[1] = maxVMerge
						mergeCellValue = text
						maxHMerge := cellIndex + h
						hMergeScope[0] = cellIndex
						hMergeScope[1] = maxHMerge
					}

					fmt.Printf("行：%d,列：%d，值:%s\n", rowIndex, cellIndex, text)
				}
			}
		}
	}
}
