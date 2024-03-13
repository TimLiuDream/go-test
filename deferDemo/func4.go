package main

func func4() int {
	i := 1
	defer func() {
		i = i + 11
	}()
	return i
}
