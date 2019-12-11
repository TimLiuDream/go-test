package main

import (
	"fmt"
	"regexp"
)

//判断文本是不是中文
func IsChinese(str string) bool {
	var result bool
	if m, _ := regexp.MatchString("^\\p{Han}+$", str); !m {
		return false
	}
	return result
}

//判断文本是不是英文
func IsEnglish(str string) bool {
	if m, _ := regexp.MatchString("^[a-zA-Z]+$", str); !m {
		return false
	}
	return true
}

//判断文本是不是数字
func IsNum(str string) bool {
	if m, _ := regexp.MatchString("^[0-9]+$", str); !m {
		return false
	}
	return true
}

//判断文本是不是邮箱地址
func IsEmail(str string) bool {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, str); !m {
		return false
	}
	return true
}

//判断文本是不是手机号码
func IsPhoneNum(str string) bool {
	if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, str); !m {
		return false
	}
	return true
}

//判断文本是否是身份证
func IsIdcard(str string) bool {
	//验证15位身份证，15位的是全部数字
	if m, _ := regexp.MatchString(`^(\d{15})$`, str); !m {
		return false
	}

	//验证18位身份证，18位前17位为数字，最后一位是校验位，可能为数字或字符X。
	if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, str); !m {
		return false
	}
	return true
}

func main() {
	func1()
}

func func1() {
	var re = regexp.MustCompile(`(?mi)(update|insert\s+into|delete\s+from)\s+\x60?([a-z_]+)?\x60?`)
	var str = `package main

	import "fmt"
	
	func main() {
		
	}
	`
	result := re.FindAllStringSubmatch(str, -1)
	fmt.Println(result)
}
