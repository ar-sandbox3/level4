package sum

// Ints sums up numbers.
func Ints(numbers ...int) int {
	return ints(numbers)
}

func ints(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return ints(numbers[1:]) + numbers[0]
}
