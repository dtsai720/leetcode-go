package src

// SingleNumberIII is a function to find two numbers that only appear once in an array.
// returns an array of two numbers that only appear once in the array.
func SingleNumberIII(nums []int) []int {
	myset := make(map[int]bool)
	for _, num := range nums {
		if _, ok := myset[num]; ok {
			delete(myset, num)
		} else {
			myset[num] = true
		}
	}

	result := make([]int, 0)
	for key := range myset {
		result = append(result, key)
	}

	return result
}
