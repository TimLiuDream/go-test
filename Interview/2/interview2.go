package main

func main() {
	pase_student()
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	// 	zhou => wang
	// li => wang
	// wang => wang
	// 	for _, stu := range stus {
	// 		m[stu.Name] = &stu
	// 	}

	// 正确
	// zhou => zhou
	// li => li
	// wang => wang
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}
	for k, v := range m {
		println(k, "=>", v.Name)
	}
}
