package solution

import (
	"math"
	"reflect"
	"sort"
)

//Solution returns all positive solutions to the equation a^3 + b^3 = c^3 + d^3 where a, b, c, d is less than 1000
func Solution() [][]int {
	var solutions [][]int
	var a, b int
	m := make(map[int][]int)

	for a = 1; a < 1000; a++ {
		for b = 1; b < 1000; b++ {
			sum := PowOfThree(a) + PowOfThree(b)
			if v, ok := m[sum]; ok == true && a+b != v[0]+v[1] {
				if set := []int{a, b, v[0], v[1]}; Found(set, solutions) == false {
					solutions = append(solutions, set)
				}
				continue
			} else {
				m[sum] = []int{a, b}
			}
		}
	}
	return solutions
}

//PowOfThree returns the number to the power of 3
func PowOfThree(num int) int {
	return int(math.Pow(float64(num), 3))
}

//Found finds out if total contains certain set
func Found(set []int, total [][]int) bool {
	var result bool

	sort.Ints(set)
	for _, s := range total {
		sort.Ints(s)
		if reflect.DeepEqual(s, set) {
			result = true
		}
	}
	return result
}
