package src

// CoinChangeII is a function to find the number of combinations that make up that amount.
// The function takes an array of distinct coin values and an integer amount. The function
// returns the number of combinations that make up that amount.
func CoinChangeII(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}
