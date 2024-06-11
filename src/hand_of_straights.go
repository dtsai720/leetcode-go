package src

import "sort"

// HandOfStraights is a function that determines if the given array of integers can be divided into k consecutive hands.
// It returns true if it can be divided, otherwise false.
func HandOfStraights(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}

	sort.Ints(nums)
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}

	for _, num := range nums {
		if count, ok := mp[num]; !ok || count == 0 {
			continue
		}

		for i := range k {
			if count, ok := mp[num+i]; !ok || count == 0 {
				return false
			}

			mp[num+i]--
		}
	}

	return true
}
