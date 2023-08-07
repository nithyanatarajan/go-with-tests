package integers

// Add takes 2 integers and returns their sum.
func Add(a, b int) int {
	return a + b
}

// Sum takes variadic integers and returns their sum.
func Sum(nums ...int) int {
	result := 0

	for _, num := range nums {
		result += num
	}

	return result
}

// SumAll takes a varying number of slices, returning a new slice containing the totals for each slice passed in

func SumAll(numsList ...[]int) []int {
	var result []int

	for _, nums := range numsList {
		result = append(result, Sum(nums...))
	}

	return result
}
