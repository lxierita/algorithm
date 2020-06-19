package solution

import (
	"math"
)

//Solution returns all positive solutions to the equation a^3 + b^3 = c^3 + d^3 where a, b, c, d is less than 1000
func Solution() [][]int {
	var solutions [][]int
	var a, b int
	m := make(map[float64][]int)
	am := make(map[int]int)

	for a = 1; a < 1000; a++ {
		for b = 1; b < 1000; b++ {
			sum := PowOfThree(a) + PowOfThree(b)
			if v, ok := m[sum]; ok == true && a+b != v[0]+v[1] {
				solutions = append(solutions, []int{a, b, v[0], v[1]})
			} else {
				m[sum] = []int{a, b}
			}
		}
	}
	for i, s := range solutions {
		sum := 0
		for _, num := range s {
			sum += num
		}
		if _, ok := am[sum]; ok == true {
			solutions = append(solutions[:i], solutions[i+1:]...)
		} else {
			am[sum] = sum
		}
	}

	return solutions
}

//PowOfThree returns the number to the power of 3
func PowOfThree(num int) float64 {
	return math.Pow(float64(num), 3)
}
