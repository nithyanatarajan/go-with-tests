package integers

// Add takes 2 integers and returns their sum.
func Add(a, b int) int {
	return a + b
}

// Sum takes variadic integers and returns their sum.
func Sum(numbersToSum ...int) int {
	result := 0

	for _, number := range numbersToSum {
		result += number
	}

	return result
}

// SumAll takes a varying number of slices,
// returning a new slice containing the totals for each slice passed in
func SumAll(numbersListToSum ...[]int) []int {
	var result []int

	for _, numbersToSum := range numbersListToSum {
		result = append(result, Sum(numbersToSum...))
	}

	return result
}

// SumAllTails takes a varying number of slices,
// returning a new slice containing the totals of the "tails" of each slice.
// The tail of a collection is all items in the collection except
// the first one (the "head").
func SumAllTails(numbersListToSum ...[]int) []int {
	var result []int

	for _, numbersToSum := range numbersListToSum {
		if len(numbersToSum) == 0 {
			result = append(result, 0)
		} else {
			// slice[low:high]
			// numbersToSum[:1] returns 1st element (upto 1)
			// numbersToSum[1:] returns all but 1st element (from 1)
			tail := numbersToSum[1:]
			result = append(result, Sum(tail...))
		}
	}
	return result
}
