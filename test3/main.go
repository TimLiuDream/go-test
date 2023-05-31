package main

func main() {

}

func GetMerge(m, n []int) []int {
	p1, p2 := 0, 0
	result := make([]int, 0)
	for {
		if p1 == len(m) || p2 == len(n) {
			break
		}
		if m[p1] == n[p2] {
			result = append(result, m[p1])
			p1++
			p2++
		} else if m[p1] < m[p2] {
			p1++
		} else {
			p2++
		}
	}
	return result
}
