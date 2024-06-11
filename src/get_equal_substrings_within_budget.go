package src

// Get the maximum length of a substring of s that the difference between
// the cost of the substring and the cost of the substring of t is less than or equal to maxCost.
// The function returns the maximum length of a substring of s that
// the difference between the cost of the substring and the cost of the substring of t is
// less than or equal to maxCost.
// returns -1 if the length of s and t are not equal.
func GetEqualSubstringsWithinBudget(s string, t string, maxCost int) int {
	if len(s) != len(t) {
		return -1
	}

	maxLength := 0
	for fast, slow := 0, -1; fast < len(s); fast++ {
		maxCost -= int(SubAbs(s[fast], t[fast]))
		for maxCost < 0 {
			slow++
			maxCost += int(SubAbs(s[slow], t[slow]))
		}

		maxLength = max(maxLength, fast-slow)
	}

	return maxLength
}
