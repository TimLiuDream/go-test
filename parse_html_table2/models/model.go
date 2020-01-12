package models

type Cell struct {
	RowIndex int
	ColIndex int
	Value    string
	MapAttr  map[string]string
}

func NewCell(rowIndex, colIndex int, value string, mapAttr map[string]string) *Cell {
	return &Cell{
		RowIndex: rowIndex,
		ColIndex: colIndex,
		Value:    value,
		MapAttr:  mapAttr,
	}
}
