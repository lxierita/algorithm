package solution

import (
	"math"
)

//Solution returns all positive solutions to the equation a^3 + b^3 = c^3 + d^3 where a, b, c, d is less than 1000
func Solution() [][]int {
	var solutions [][]int
	var a, b int
	m := make(map[float64][]int)

	for a = 1; a < 1000; a++ {
		for b = 1; b < 1000; b++ {
			sum := PowOfThree(a) + PowOfThree(b)
			if v, ok := m[sum]; ok == true && a+b != v[0]+v[1] {
				if set := []int{a, b, v[0], v[1]}; found(set, solutions) == false {
					solutions = append(solutions, set)
				}

			} else {
				m[sum] = []int{a, b}
			}
		}
	}
	return solutions
}

//PowOfThree returns the number to the power of 3
func PowOfThree(num int) float64 {
	return math.Pow(float64(num), 3)
}

func found(set []int, total [][]int) bool {
	var result bool
	setsum := 0

	for _, y := range set {
		setsum += y
	}

	for _, s := range total {
		sum := 0
		for _, x := range s {
			sum += x
		}
		if sum == setsum {
			result = true
		}
	}
	return result
}
