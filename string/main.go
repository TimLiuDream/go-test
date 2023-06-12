package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/axgle/mahonia"
)

func main() {
	func11()
}

func func1() {
	s := "hello"
	s = "c" + s[1:]
	fmt.Printf("%s\n", s)
}

func func2() {
	str := "lkdjfl000"
	bytes := []rune(str)
	for from, to := 0, len(bytes)-1; from < to; from, to = from+1, to-1 {
		bytes[from], bytes[to] = bytes[to], bytes[from]
	}
	str = string(bytes)
	fmt.Print(str)
}

func func3() {
	var s *string
	s = new(string)
	fmt.Println(s)
}

func func4() {
	str := "乱码的字符串变量"
	str = ConvertToString(str, "gbk", "utf-8")
	fmt.Println(str)
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func func5() {
	str := strings.Replace("ssfsdfsdfsdfs", "dev", "single", -1)
	fmt.Println(str)
}

func func6() {
	str := `abc` + `123`
	fmt.Println(str)
}

func func7() {
	teamUUIDs := []string{"123", "345"}
	apps := []string{"wps", "automation"}
	StatusEnable := 1 // 待安装

	now := time.Now().Unix()
	insertAppInstalledSql := "'use project;insert into appstore_app_installed values %s;'"
	values := strings.Builder{}
	for _, teamUUID := range teamUUIDs {
		for _, app := range apps {
			values.WriteString("('")
			values.WriteString(teamUUID)
			values.WriteString("','")
			values.WriteString(app)
			values.WriteString("',")
			values.WriteString(strconv.FormatInt(now, 10))
			values.WriteString(",")
			values.WriteString(strconv.FormatInt(int64(StatusEnable), 10))
			values.WriteString(")")
			values.WriteString(",")
		}
	}
	valuesStr := values.String()
	valuesStr = valuesStr[:len(valuesStr)-1]
	insertAppInstalledSql = fmt.Sprintf(insertAppInstalledSql, valuesStr)
	fmt.Println(insertAppInstalledSql)
}

type stringAlias = string

func func10() {
	var s1 stringAlias = "2"
	v, ok := interface{}(s1).(string)
	fmt.Println(ok)
	fmt.Println(v)
	fmt.Println(reflect.TypeOf(s1))
}

type stringAlias1 string

func func11() {
	var s1 stringAlias1 = "2"
	v, ok := interface{}(s1).(string)
	fmt.Println(ok)
	fmt.Println(v)
	fmt.Println(reflect.TypeOf(s1))
}

func PreAllocStringsBuilder(n int, str string) string {
	var builder strings.Builder

	//预分配
	builder.Grow(n * len(str))

	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}
