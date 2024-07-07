package main

// 문제:
// 분류:
func solution(lottos []int, win_nums []int) []int {
	set := make(map[int]struct{})
	for _, num := range win_nums {
		set[num] = struct{}{}
	}

	zeros := 0
	match := 0

	for _, num := range lottos {
		if num == 0 {
			zeros++
			continue
		}

		if _, ok := set[num]; ok {
			match++
		}
	}

	return []int{
		7 - max(match+zeros, 1),
		7 - max(match, 1),
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
}
