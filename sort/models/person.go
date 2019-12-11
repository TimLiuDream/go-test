package models

type Person struct {
	Name string
	Age  int
}

type PersonSlice []Person

func (s PersonSlice) Len() int           { return len(s) }
func (s PersonSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s PersonSlice) Less(i, j int) bool { return s[i].Age < s[j].Age }
