package slice

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("colection with 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Sum(numbers)
		expect := 15
		if expect != result {
			t.Errorf("expect '%d' - result '%d' - data '%v'", expect, result, numbers)
		}
	})
}

func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		_ = Sum(numbers)
	}
}
