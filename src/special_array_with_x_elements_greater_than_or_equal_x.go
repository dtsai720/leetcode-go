package src

import (
	"sort"
)

// SpecialArrayWithXElementsGreaterThanOrEqualX returns the special array with x elements greater than or equal to x.
// nums is a non-empty array of integers.
// The special array is an array where the number of elements is equal to x and all elements are greater than or equal to x.
// If there is no such x, return -1.
//
//nolint:mnd
func SpecialArrayWithXElementsGreaterThanOrEqualX(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	sort.Ints(nums)
	minValue, maxValue := 0, nums[len(nums)-1]
	for minValue <= maxValue {
		midValue := minValue + (maxValue-minValue)/2
		index := LowerBound(nums, midValue)
		if len(nums)-index == midValue {
			return midValue
		}
		if len(nums)-index > midValue {
			minValue = midValue + 1
		} else {
			maxValue = midValue - 1
		}
	}

	return -1
}
