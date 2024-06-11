package src

// LowerBound is a function to find the lower bound of a number in a sorted array.
// The function returns the index of the number in the array.
//
//nolint:mnd
func LowerBound(array []int, num int) int {
	l, r := 0, len(array)
	for l < r {
		m := l + (r-l)/2
		if array[m] < num {
			l = m + 1
		} else {
			r = m
		}
	}

	return l
}

// The function returns the absolute difference between the two numbers.
//
//nolint:ireturn
func SubAbs[T byte | int](a, b T) T {
	if a > b {
		return a - b
	}

	return b - a
}
