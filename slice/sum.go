package slice

// Sum receives an array of 5 numbers and resumes their sum
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
