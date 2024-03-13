package main

func func5() (r *int) {
	r = new(int)
	defer func(r *int) {
		*r = *r + 3
	}(r)
	*r = 2
	return r
}
