package main

import "fmt"

func StringArrayDifference(old []string, new []string) (additions []string, deletions []string) {
	additionsMap := make(map[string]struct{})
	deletionsMap := make(map[string]struct{})
	for _, s := range old {
		if len(s) > 0 {
			deletionsMap[s] = struct{}{}
		}
	}
	for _, s := range new {
		if len(s) > 0 {
			additionsMap[s] = struct{}{}
		}
	}
	for s, _ := range additionsMap {
		if _, ok := deletionsMap[s]; !ok {
			additions = append(additions, s)
		}
	}
	for s, _ := range deletionsMap {
		if _, ok := additionsMap[s]; !ok {
			deletions = append(deletions, s)
		}
	}
	return
}

func main() {
	s1 := []string{"1", "2", "3"}
	s2 := []string{"1", "2", "3"}

	additions, deletions := StringArrayDifference(s1, s2)
	fmt.Println(additions)
	fmt.Println(deletions)
}
