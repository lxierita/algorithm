package solution_test

import (
	"fmt"
	. "solution"
	"testing"
)

func TestSolution_should_return_all_positive_integer_solutions(t *testing.T) {
	got := Solution()
	fmt.Println(got)
	for _, set := range got {
		m := make(map[int]float64)
		for i, n := range set {
			if n > 1000 || n < 1 {
				t.Errorf("a, b, c, d should be less than 1000, got %d", n)
			}
			m[i] = PowOfThree(n)
		}
		if m[0]+m[1] != m[2]+m[3] {
			t.Errorf("%v + %v != %v + %v", m[0], m[1], m[2], m[3])
		}

	}
}
