package main

import (
	"fmt"
	"sort"

	utilsModel "github.com/bangwork/bang-api/app/models/utils"
	"github.com/bangwork/bang-api/app/utils"
)

type Test struct {
	Name string
}

type TestSorter []*Test

func (t TestSorter) Len() int      { return len(t) }
func (t TestSorter) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t TestSorter) Less(i, j int) bool {
	return utils.CompareString(utilsModel.TranslationPinyin(t[i].Name),
		utilsModel.TranslationPinyin(t[j].Name))
}

func main() {
	tests := make([]*Test, 0)
	test1 := &Test{
		Name: "b",
	}
	test2 := &Test{
		Name: "B",
	}
	test3 := &Test{
		Name: "z",
	}
	test4 := &Test{
		Name: "b",
	}
	tests = append(tests, test1)
	tests = append(tests, test2)
	tests = append(tests, test3)
	tests = append(tests, test4)
	sort.Sort(TestSorter(tests))
	for _, test := range tests {
		fmt.Println(test.Name)
	}
}
