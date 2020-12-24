package main

import "fmt"

func main() {
	s := "{\"condition_groups\":[[{\"field_uuid\":\"field024\",\"operate\":{\"label\":\"filter.addQueryContent.include\",\"operate_id\":\"include\",\"predicate\":\"in\"},\"value\":[\"%s\",\"%s\"]}]]}"
	placeholder := []interface{}{"11111", "222222"}
	fmt.Println(fmt.Sprintf(s, placeholder...))
}
