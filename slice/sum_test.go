package slice

import (
	"reflect"
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
func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2}, []int{0, 9})
	expect := []int{3, 9}
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("expect '%v' - result '%v'", expect, result)
	}
}

func TestSumAllRest(t *testing.T) {
	verifySums := func(t *testing.T, result, expect []int) {
		t.Helper()
		if !reflect.DeepEqual(result, expect) {
			t.Errorf("expect '%v' - result '%v'", expect, result)
		}
	}

	t.Run("makes the sums of some slices", func(t *testing.T) {
		result := SumAllRest([]int{1, 2}, []int{0, 9})
		expect := []int{2, 9}
		verifySums(t, result, expect)
	})

	t.Run("sum empty slices securely", func(t *testing.T) {
		result := SumAllRest([]int{}, []int{3, 4, 5})
		expect := []int{0, 9}
		verifySums(t, result, expect)
	})
}

func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		_ = Sum(numbers)
	}
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SumAll([]int{1, 2}, []int{0, 9})
	}
}
