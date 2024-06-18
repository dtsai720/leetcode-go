package src

import "sort"

func MostProfitAssigningWork(difficulty []int, profit []int, worker []int) int {
	n := len(difficulty)
	jobs := make([][2]int, 0, n)
	for i := range n {
		jobs = append(jobs, [2]int{difficulty[i], profit[i]})
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][0] < jobs[j][0]
	})

	sort.Ints(worker)
	result, maxProfit := 0, 0
	idx := 0

	for _, w := range worker {
		for ; idx < n && jobs[idx][0] <= w; idx++ {
			maxProfit = max(maxProfit, jobs[idx][1])
		}

		result += maxProfit
	}

	return result
}
