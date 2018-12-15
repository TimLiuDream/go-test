package Model

type Test struct {
	Id int
	Year int
	Amount int
	Num int
}

func (Test) TableName() string {
	return "test"
}