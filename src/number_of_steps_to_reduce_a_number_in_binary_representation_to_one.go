package src

// NumberOfStepsToReduceANumberInBinaryRepresentationToOne ...
// Given a number s in binary representation, you can always reduce it to 1 by performing the following steps:
// If the number is even, divide it by 2.
// If the number is odd, add 1 to it.
// A binary representation is a string consisting of only 0's and 1's.
// Return the number of steps to reduce s to 1.
func NumberOfStepsToReduceANumberInBinaryRepresentationToOne(s string) int {
	count := 0
	var carry byte = 0
	for i := len(s) - 1; i > 0; i-- {
		if (s[i]-'0'+carry)%2 == 1 {
			count += 2
			carry = 1
		} else {
			count++
		}
	}

	return count + int(carry)
}
