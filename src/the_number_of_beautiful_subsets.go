package src

import (
	"slices"
	"sort"
)

// You are given an array nums of positive integers and a positive integer k.
// A subset of nums is beautiful if it does not contain two integers
// with an absolute difference equal to k.
// Return the number of non-empty beautiful subsets of the array nums.
// A subset of nums is an array that can be obtained by deleting some (possibly none) elements
// from nums. Two subsets are different if and only if the chosen indices to delete are different.
// The answer may be very large, return it modulo 10^9 + 7.
func TheNumberOfBeautifulSubset(nums []int, k int) int {
	sort.Ints(nums)

	var f func(idx int, array []int) int

	f = func(idx int, array []int) int {
		if idx == len(nums) {
			return 0
		}

		count := f(idx+1, array)
		if _, ok := slices.BinarySearch(array, nums[idx]-k); !ok {
			count += 1 + f(idx+1, append(array, nums[idx]))
		}

		return count
	}

	return f(0, nil)
}
