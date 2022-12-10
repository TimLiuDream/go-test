package collections

import (
	"fmt"
	"testing"
)

func TestSets(t *testing.T) {
	A := SliceToSet([]int{1, 3, 5})
	B := SliceToSet([]int{0, 1, 2, 3, 4})

	union := A.Union(B)
	fmt.Println(union) // map[0:true 1:true 2:true 3:true 4:true 5:true]

	C := SliceToSet([]string{"a", "b", "noah"})
	D := SliceToSet([]string{"a", "noah"})
	intersection := C.Intersection(D)
	fmt.Println(intersection) // map[a:true noah:true]

	fmt.Println(C.Equals(D)) // false
}
