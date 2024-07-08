package main

import "fmt"

// 문제: https://school.programmers.co.kr/learn/courses/30/lessons/12946
// 분류: 재귀, 하노이의 탑
func solution(n int) [][]int {
	// 하노이의 탑
	// n개의 원판을 1번 기둥에서 3번 기둥으로 옮기는 방법을 구하라
	// 단, 한 번에 한 개의 원판만 옮길 수 있다.
	// 원판은 항상 큰 원판이 작은 원판 위에 있어서는 안된다.

	answer := [][]int{}

	var hanoi func(n, from, to, via int)

	hanoi = func(n, from, to, via int) {
		// n이 1이면 from에서 to로 옮기고 종료
		if n == 1 {
			answer = append(answer, []int{from, to})
			return
		}

		hanoi(n-1, from, via, to)                // n-1개의 원판을 from에서 to를 거쳐 via로 옮긴다.
		answer = append(answer, []int{from, to}) // n번째 원판을 from에서 to로 옮긴다.
		hanoi(n-1, via, to, from)                // n-1개의 원판을 via에서 from을 거쳐 to로 옮긴다.
	}

	hanoi(n, 1, 3, 2) // n개의 원판을 1번 기둥에서 3번 기둥으로 옮긴다. (2번 기둥을 거쳐서)

	// 최소 이동 횟수: 2^n - 1

	return answer
}

func main() {
	answer := solution(3)
	fmt.Println(answer)
}
