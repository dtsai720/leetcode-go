package src

// StudentAttendanceRecordII is a function to calculate the number of possible attendance records
// for a student with n days of attendance.
// The attendance record may contain the following:
// 1. 'A' for absent
// 2. 'L' for late
// 3. 'P' for present
// The student is eligible for an award if the attendance record doesn't contain more than one 'A' and
// doesn't contain three consecutive 'L's.
// The function returns the number of possible attendance records.
// returns -1 if n is less than 1.
func StudentAttendanceRecordII(n int) int {
	const (
		MaxLate   = 3
		MaxAbsent = 2
		MOD       = 1e9 + 7
	)

	dp := make([][][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([][]int, MaxAbsent+1)
		for j := 0; j <= 2; j++ {
			dp[i][j] = make([]int, MaxLate+1)
			for k := 0; k <= 3; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	var f func(n, absent, late int) int

	f = func(n, absent, late int) int {
		if absent == 2 || late == 3 {
			return 0
		}

		if n == 0 {
			return 1
		}

		if dp[n][absent][late] != -1 {
			return dp[n][absent][late]
		}

		output := f(n-1, absent, 0) % MOD
		output = (output + (f(n-1, absent+1, 0) % MOD)) % MOD
		output = (output + (f(n-1, absent, late+1) % MOD)) % MOD
		dp[n][absent][late] = output

		return output
	}

	return f(n, 0, 0)
}
