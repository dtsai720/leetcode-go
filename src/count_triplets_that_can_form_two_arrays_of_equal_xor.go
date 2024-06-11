package src

// CountTripletsThatCanFormTwoArraysOfEqualXor is a function to count the number of triplets
// that can form two arrays of equal XOR. The function takes an array of integers and returns
// the number of triplets that can form two arrays of equal XOR.
func CountTripletsThatCanFormTwoArraysOfEqualXor(nums []int) int {
	cnt, total := make(map[int]int), make(map[int]int)
	cnt[0] = 1
	output, current := 0, 0

	for idx, num := range nums {
		current ^= num
		if count, ok := cnt[current]; ok {
			output += count*idx - total[current]
		}

		cnt[current]++
		total[current] += idx + 1
	}

	return output
}
