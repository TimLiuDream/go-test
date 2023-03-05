package generic_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func filterStrings(collection []string, test func(string) bool) (f []string) {
	for _, s := range collection {
		if test(s) {
			f = append(f, s)
		}
	}
	return
}

func filter[T any](collection []T, test func(T) bool) (f []T) {
	for _, e := range collection {
		if test(e) {
			f = append(f, e)
		}
	}

	return
}

func generateStringCollection(size, strLen int) []string {
	var collection []string
	for i := 0; i < size; i++ {
		collection = append(collection, strings.Repeat(fmt.Sprintf("%c", rune('A'+(i%10))), strLen))
	}

	return collection
}

func TestFilter(t *testing.T) {
	c := generateStringCollection(1000, 3)

	t.Run("the output of the typed and generic functions is the same", func(t *testing.T) {
		predicate := func(s string) bool { return s == "AAA" }

		filteredStrings := filterStrings(c, predicate)
		filteredElements := filter(c, predicate)

		if !reflect.DeepEqual(filteredStrings, filteredElements) {
			t.Errorf("the output of the two functions is not the same")
		}
	})
}
