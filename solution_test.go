package solution_test

import (
	"fmt"
	. "solution"
	"testing"
)

func TestSolution_should_return_all_positive_integer_solutions(t *testing.T) {
	got := Solution()
	fmt.Println(len(got))
	for _, set := range got {
		m := make(map[int]int)
		// fmt.Println(set)
		for i, n := range set {
			if n > 1000 || n < 1 {
				t.Errorf("a, b, c, d should be less than 1000, got %d", n)
			}
			m[i] = PowOfThree(n)
		}
		// if m[0]+m[1] != m[2]+m[3] {
		// 	t.Errorf("%v + %v != %v + %v", m[0], m[1], m[2], m[3])
		// }
	}
}

func TestPowOfThree_returns_num_to_the_power_of_three(t *testing.T) {
	num1 := 1
	num2 := 2
	if got := PowOfThree(num1); got != 1 {
		t.Errorf("PowOfThree(%v) = %v; should return 1", num1, got)
	}
	if got := PowOfThree(num2); got != 8 {
		t.Errorf("PowOfThree(%v) = %v; should return 8", num2, got)
	}
}

func TestFound_finds_out_if_the_same_set_of_different_order_exists_in_collection(t *testing.T) {

	set1 := []int{1, 2, 3}
	set2 := []int{67, 34, 23}
	total := [][]int{{2, 3, 4}, {1, 3, 2}, {4, 5, 6}}
	if got := Found(set1, total); got != true {
		t.Errorf("Found() = %v; should be true", got)
	}
	if got := Found(set2, total); got != false {
		t.Errorf("Found() = %v; should be false", got)
	}

}
