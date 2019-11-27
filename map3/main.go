package main

import "fmt"

func main() {
	oldMap := map[string]string{"1": "1", "2": "2", "3": "3"}
	newFields := []string{"1", "2", "4"}
	for _, field := range newFields {
		if _, found := oldMap[field]; found {
			delete(oldMap, field)
		}
	}
	if len(oldMap) > 0 {
		for key := range oldMap {
			newFields = append(newFields, key)
		}
	}
	fmt.Println(newFields)
}
