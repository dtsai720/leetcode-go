package src

func SlidingPuzzle(board [][]int) int {
	const target = "123450"

	array := [][]int{
		{1, 3},
		{0, 2, 4},
		{1, 5},
		{0, 4},
		{1, 3, 5},
		{2, 4},
	}

	ms := make(map[string]bool)

	var start []rune
	for _, nums := range board {
		for _, num := range nums {
			start = append(start, rune(num)+'0')
		}
	}

	queue := []string{string(start)}
	ms[string(start)] = true
	count := 0
	for len(queue) > 0 {
		size := len(queue)
		for range size {
			cur := queue[0]
			queue = queue[1:]
			if cur == target {
				return count
			}

			zero := 0
			for cur[zero] != '0' {
				zero++
			}

			for _, next := range array[zero] {
				nextBoard := []rune(cur)
				nextBoard[zero], nextBoard[next] = nextBoard[next], nextBoard[zero]
				nextStr := string(nextBoard)
				if _, ok := ms[nextStr]; ok {
					continue
				}
				ms[nextStr] = true
				queue = append(queue, nextStr)
			}
		}

		count++
	}

	return -1
}
