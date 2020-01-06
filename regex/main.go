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

func IsDomainName(domainName string) bool {
	domainNameExpr := `^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,6}$`
	domainNameCompile := regexp.MustCompile(domainNameExpr)
	return domainNameCompile.MatchString(domainName)
}

func IsUrl(url string) bool {
	urlExpr := `[a-zA-z]+://[^\s]*`
	urlCompile := regexp.MustCompile(urlExpr)
	return urlCompile.MatchString(url)
}

func main() {
	// func1()
	fmt.Println(IsUrl("https://www.baidu.com"))
	fmt.Println(IsUrl("https://www.baidu.com/s?wd=%E5%88%87%E6%8D%A2python%E8%99%9A%E6%8B%9F%E7%8E%AF%E5%A2%83%E4%BB%A3%E7%A0%81\u0026rsv_spt=1\u0026rsv_iqid=0xe04a453e0001d8cb\u0026issp=1\u0026f=8\u0026rsv_bp=1\u0026rsv_idx=2\u0026ie=utf-8\u0026tn=baiduhome_pg\u0026rsv_enter=1\u0026rsv_dl=tb\u0026rsv_sug3=44\u0026rsv_sug1=28\u0026rsv_sug7=101\u0026rsv_t=05700IcHf6K0d1FATNPlII%2Fc0jBDV3tkgj%2F3rl0uXvOtk4EdT99GqRx%2F2nU6lML3s3B3\u0026rsv_sug2=0\u0026inputT=16006\u0026rsv_sug4=16583"))
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
