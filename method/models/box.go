package models

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	Width, Height, Depth float64
	Color                Color
}

type BoxList []Box

func (b Box) Volume() float64 {
	return b.Width * b.Height * b.Depth
}

func (b *Box) SetColor(c Color) {
	b.Color = c
}

func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)

	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.Color
		}
	}
	return k
}

func (bl BoxList) PaintItBlack() {
	for i, _ := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}
