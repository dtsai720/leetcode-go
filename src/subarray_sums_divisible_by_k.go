package src

// SubarraySumsDivisibleByK returns the number of subarrays with sum divisible by k.
// The function takes an array of integers and an integer k.
// The function returns the number of subarrays.
func SubarraySumsDivisibleByK(nums []int, k int) int {
	count := make([]int, k)
	count[0] = 1
	sum, output := 0, 0
	for _, num := range nums {
		sum = ((sum+num)%k + k) % k
		output += count[sum]
		count[sum]++
	}

	return output
}
