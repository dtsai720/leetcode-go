package src

import "errors"

var ErrLengthMismatch = errors.New("students and mentors length are not equal")

// MaximumCompatibilityScoreSumSolution is the solution for the Maximum Compatibility Score Sum problem.
// The solution uses DFS to find the maximum compatibility score sum.
// The solution has the following fields:
// - array: a 2D array that stores the compatibility score of students and mentors.
// - visited: a boolean array that stores the visited status of mentors.
// - m: the number of students.
// - n: the number of mentors.
// - Result: the maximum compatibility score sum.
type MaximumCompatibilityScoreSumSolution struct {
	array   [][]int
	visited []bool
	m, n    int
	Result  int
}

// Init initializes the solution with the given students and mentors.
// The function returns an error if the length of students and mentors are not equal.
func (mc *MaximumCompatibilityScoreSumSolution) Init(students [][]int, mentors [][]int) error {
	if len(students) != len(mentors) {
		return ErrLengthMismatch
	}

	mc.n = len(students)
	mc.m = len(students[0])

	mc.array = make([][]int, mc.n)
	for i := range mc.n {
		mc.array[i] = make([]int, mc.n)
	}

	for x := range mc.n {
		for y := range mc.n {
			if len(students[x]) != len(mentors[y]) || len(students[x]) != mc.m {
				return ErrLengthMismatch
			}

			for z := range mc.m {
				if students[x][z] == mentors[y][z] {
					mc.array[x][y]++
				}
			}
		}
	}

	mc.visited = make([]bool, mc.n)

	return nil
}

// DFS is a helper function that performs DFS to find the maximum compatibility score sum.
// The function takes two parameters, x and sum, where x is the current student index and sum is the current compatibility score sum.
// The function updates the Result field with the maximum compatibility score sum.
func (mc *MaximumCompatibilityScoreSumSolution) DFS(x, sum int) {
	if x == mc.n {
		mc.Result = max(mc.Result, sum)

		return
	}

	for i := range mc.n {
		if mc.visited[i] {
			continue
		}

		mc.visited[i] = true
		mc.DFS(x+1, sum+mc.array[x][i])
		mc.visited[i] = false
	}
}

// MaximumCompatibilityScoreSum returns the maximum compatibility score sum.
// The function takes two 2D arrays, students and mentors, where each array
// represents the compatibility of students and mentors. The function returns
// the maximum compatibility score sum.
// returns -1 if the length of students and mentors are not equal.
func MaximumCompatibilityScoreSum(students [][]int, mentors [][]int) int {
	var solution MaximumCompatibilityScoreSumSolution
	if err := solution.Init(students, mentors); err != nil {
		return -1
	}

	solution.DFS(0, 0)

	return solution.Result
}
