package main

func main() {

	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
}
//函数返回值名字会在函数起始处被初始化为对应类型的零值并且作用域为整个函数
//4
func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

//t作用于函数内部，但是除了defer内部
//1
func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}
//函数返回值名字会在函数起始处被初始化为对应类型的零值并且作用域为整个函数
//3
func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}