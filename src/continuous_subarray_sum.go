package src

// ContinuousSubarraySum returns true if there is a continuous subarray that sums up to k.
// The function takes an array of integers and an integer k. The function returns true if
// there is a continuous subarray that sums up to k.
func ContinuousSubarraySum(nums []int, k int) bool {
	var sum int64 = 0
	seen := make(map[int64]bool)
	m := int64(k)
	for _, num := range nums {
		sum = (sum + int64(num)) % m
		if _, ok := seen[sum]; ok {
			return true
		}

		seen[((sum-int64(num))%m+m)%m] = true
	}

	return false
}
