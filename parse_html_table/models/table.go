package models

type Cell struct {
	RowIndex int
	ColIndex int
	Value    string
}

func AddCell(mapRowCells map[int][]*Cell, rowIndex int, colIndex int, value string) *Cell {
	cells := mapRowCells[rowIndex]
	for _, cell := range cells {
		if cell.ColIndex == colIndex {
			colIndex++
		}
	}
	c := &Cell{
		RowIndex: rowIndex,
		ColIndex: colIndex,
		Value:    value,
	}
	return c
}

type CellSorter []*Cell

func (t CellSorter) Len() int      { return len(t) }
func (t CellSorter) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t CellSorter) Less(i, j int) bool {
	if t[i].RowIndex > t[j].RowIndex {
		return false
	} else if t[i].RowIndex < t[j].RowIndex {
		return true
	} else {
		return t[i].ColIndex < t[j].ColIndex
	}
}
