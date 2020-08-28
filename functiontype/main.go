package main

import "fmt"

type RuleVarify func() bool

func (rv RuleVarify) verify() bool {
	return rv()
}

func CreateTaskRuleV1Varify() bool {
	return false
}

func CreateTaskRuleV2Varify() bool {
	return true
}

func main() {
	rv := RuleVarify(CreateTaskRuleV1Varify)
	fmt.Println(rv.verify())
	rv = RuleVarify(CreateTaskRuleV2Varify)
	fmt.Println(rv.verify())
}
