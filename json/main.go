package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	// "github.com/chanxuehong/util/json"
)

// 这里的字段需要注意
// 当字段是小写字母开头的话，是不可导出的，所以在序列化的时候并不能正确的序列化,即使有后面的标记也是不可以的
// 字段一定要大写开头
type Result struct {
	// status int `json:"status"`
	Status int
}

type Seq struct {
	Num  int
	Name string
}
type Info struct {
	Id    int
	Value string
}
type firstCase struct {
	*Seq
	*Info
}

func main() {
	// func1()
	// func3()
	// func4()
	func5()
}

func func1() {
	var s string = `{"status":200}`
	r := &Result{}

	err := json.Unmarshal([]byte(s), r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(r)
	fmt.Println(r.Status)
	// nA := Seq{1, "11"}
	// nB := Info{101, "aaa"}
	// var nF = firstCase{&nA, &nB}
	// data, _ := json.Marshal(nF)
	// fmt.Println(data)
	// fmt.Println(string(data))
}

func func2() {
	nA := Seq{1, "11"}
	nB := Info{101, "aaa"}
	//方法1
	anyMarshal(struct {
		*Seq
		*Info
	}{&nA, &nB})
	//方法2
	anyMarshalNew(nA, nB)
}

func anyMarshal(msg interface{}) {
	mData, _ := json.Marshal(msg)
	fmt.Println(mData)
	fmt.Println(string(mData))
}

func anyMarshalNew(nA Seq, nB interface{}) {
	nF := struct {
		*Seq
		Any interface{}
	}{&nA, &nB}
	mData, _ := json.Marshal(nF)
	fmt.Println(mData)
	fmt.Println(string(mData))
}

type test struct {
	Changelogs string `json:"changelogs"`
}

func func3() {
	str := `{
		"version_code": 4840,
		"version_name": "2.22.0",
		"forcibly": true,
		"url": "https://ones.ai/downloads/ones_release_v2.22.0_build4840_huafa.apk",
		"size": 14416886,
		"sha1": "63702f00dbbd2773d937f8df4d05f16ba530ca83",
		"changelogs": "主要更新:
		- 修复缺陷，提高稳定性。
		",
		"modify_time": 1584786131
	}`

	test := new(test)
	err := json.Unmarshal([]byte(str), &test)
	if err != nil {
		str = strings.Replace(str, "\n", "", -1)
		// str = strings.Replace(str, "\t", "", -1)
		err = json.Unmarshal([]byte(str), &test)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(test.Changelogs)
}

func func4() {
	str := `
	{
		"activity_chart": {
			"uuid_index": 0,
			"name_index": 3
		},
		"activity_draft": {
			"uuid_index": 0,
			"name_index": 5
		},
		"card": {
			"uuid_index": 0,
			"name_index": 3
		},
		"component": {
			"uuid_index": 0,
			"name_index": 5
		},
		"component_template": {
			"uuid_index": 0,
			"name_index": 2
		},
		"component_view": {
			"uuid_index": 0,
			"name_index": 4
		},
		"component_view_user": {
			"uuid_index": 0,
			"name_index": 5
		},
		"dashboard": {
			"uuid_index": 0,
			"name_index": 3
		},
		"deliverable": {
			"uuid_index": 0,
			"name_index": 3
		},
		"department": {
			"uuid_index": 0,
			"name_index": 2
		},
		"field": {
			"uuid_index": 0,
			"name_index": 3
		},
		"field_option": {
			"uuid_index": 0,
			"name_index": 3
		},
		"filter": {
			"uuid_index": 0,
			"name_index": 3
		},
		"gantt_chart": {
			"uuid_index": 0,
			"name_index": 2
		},
		"gantt_data": {
			"uuid_index": 0,
			"name_index": 4
		},
		"issue_type": {
			"uuid_index": 0,
			"name_index": 2
		},
		"manhour_report": {
			"uuid_index": 0,
			"name_index": 3
		},
		"milestone": {
			"uuid_index": 0,
			"name_index": 5
		},
		"organization": {
			"uuid_index": 0,
			"name_index": 3
		},
		"pipeline": {
			"uuid_index": 0,
			"name_index": 3
		},
		"plugin": {
			"uuid_index": 0,
			"name_index": 1
		},
		"ppm_task": {
			"uuid_index": 0,
			"name_index": 5
		},
		"product": {
			"uuid_index": 0,
			"name_index": 2
		},
		"product_component": {
			"uuid_index": 0,
			"name_index": 11
		},
		"product_component_view": {
			"uuid_index": 0,
			"name_index": 4
		},
		"product_module": {
			"uuid_index": 0,
			"name_index": 3
		},
		"program": {
			"uuid_index": 0,
			"name_index": 4
		},
		"program_role": {
			"uuid_index": 0,
			"name_index": 4
		},
		"project": {
			"uuid_index": 0,
			"name_index": 1
		},
		"project_field": {
			"uuid_index": 0,
			"name_index": 6
		},
		"project_filter": {
			"uuid_index": 0,
			"name_index": 4
		},
		"project_report": {
			"uuid_index": 0,
			"name_index": 3
		},
		"project_sprint_field": {
			"uuid_index": 0,
			"name_index": 4
		},
		"project_sprint_field_option": {
			"uuid_index": 0,
			"name_index": 3
		},
		"project_sprint_status": {
			"uuid_index": 0,
			"name_index": 3
		},
		"resource": {
			"uuid_index": 0,
			"name_index": 9
		},
		"role": {
			"uuid_index": 0,
			"name_index": 2
		},
		"sprint": {
			"uuid_index": 0,
			"name_index": 3
		},
		"task": {
			"uuid_index": 0,
			"name_index": 13
		},
		"task_gantt_chart": {
			"uuid_index": 0,
			"name_index": 3
		},
		"task_gantt_data": {
			"uuid_index": 0,
			"name_index": 4
		},
		"task_link_type": {
			"uuid_index": 0,
			"name_index": 3
		},
		"task_status": {
			"uuid_index": 0,
			"name_index": 2
		},
		"team": {
			"uuid_index": 0,
			"name_index": 3
		},
		"team_report": {
			"uuid_index": 0,
			"name_index": 2
		},
		"testcase_case": {
			"uuid_index": 0,
			"name_index": 6
		},
		"testcase_case_step": {
			"uuid_index": 0,
			"name_index": 5
		},
		"testcase_field": {
			"uuid_index": 0,
			"name_index": 7
		},
		"testcase_field_config": {
			"uuid_index": 0,
			"name_index": 2
		},
		"testcase_library": {
			"uuid_index": 0,
			"name_index": 2
		},
		"testcase_module": {
			"uuid_index": 0,
			"name_index": 5
		},
		"testcase_plan": {
			"uuid_index": 0,
			"name_index": 3
		},
		"transition": {
			"uuid_index": 0,
			"name_index": 6
		},
		"user": {
			"uuid_index": 0,
			"name_index": 5
		},
		"user_filter_view": {
			"uuid_index": 0,
			"name_index": 4
		},
		"user_group": {
			"uuid_index": 0,
			"name_index": 3
		},
		"version": {
			"uuid_index": 0,
			"name_index": 3
		}
	}
	`
	m := make(map[string]map[string]int)
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		log.Fatal(err)
	}
}

func func5() {
	var data []byte
	us := make([]string, 0)
	e := json.Unmarshal(data, &us)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("done")
	}
}
