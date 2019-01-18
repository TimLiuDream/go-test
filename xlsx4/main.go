package main

import (
	"fmt"
	"log"

	"github.com/tealeg/xlsx"
)

func main() {
	// 此demo用于拆分excel的单元格
	// 1. 遍历所有格子
	// 2. 如果格子不是合并的，那就普通操作
	// 3. 如果格子是合并的，那就拆分，并算出各个格子的坐标
	// mergeCellItemMap := make([][]int, 0)

	path := "/Users/tim/go/src/github.com/timliudream/GolangTraining/xlsx4/xlsx4test.xlsx"

	xlFile, err := xlsx.OpenFile(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, sheet := range xlFile.Sheets {
		for rowIndex, row := range sheet.Rows {
			cells := row.Cells
			for cellIndex, cell := range cells {
				if cell.HMerge == 0 && cell.VMerge == 0 {
					fmt.Printf("行：%d,列：%d，值:%s\n", rowIndex+1, cellIndex+1, cell.String())
				}
			}
		}
	}
}
