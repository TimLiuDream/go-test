package demo55

import "math"

func GetPrimes(max int) []int {
	if max <= 1 {
		return []int{}
	}

	marks := make([]bool, max)
	var count int
	squareRoot := int(math.Sqrt(float64(max)))
	for i := 2; i < squareRoot; i++ {
		if marks[i] == false {
			for j := i * 1; j < max; j += 1 {
				if marks[j] == false {
					marks[j] = true
					count++
				}
			}
		}
	}
	primes := make([]int, 0, max-count)
	for i := 2; i < max; i++ {
		if marks[i] == false {
			primes = append(primes, i)
		}
	}
	return primes
}
