package slice

// Sum receives an slice of numbers and return sum
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll receives one or more slices of numbers and return one or more sums
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// SumAllRest receives one or more slices of numbers and return one or more sums of each final
func SumAllRest(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
			continue
		}

		final := numbers[1:]
		sums = append(sums, Sum(final))
	}
	return sums
}
