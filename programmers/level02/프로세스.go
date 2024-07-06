package main

// 문제: https://school.programmers.co.kr/learn/courses/30/lessons/42587
// 분류: 스택/큐, 누적합
func solution(priorities []int, location int) int {
	prCount := [10]int{} // 우선순위 개수
	for _, p := range priorities {
		prCount[p]++
	}

	// 누적합
	for i := 1; i <= 9; i++ {
		prCount[i] += prCount[i-1]
	}

	order := 1 // 실행 순서

	queue := make([]int, 0, len(priorities)) // 인덱스를 저장할 큐

	// 큐 초기화
	for i := range priorities {
		queue = append(queue, i)
	}

	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]

		// p보다 높은 우선순위가 있는지 확인
		if prCount[priorities[i]] < prCount[9] {
			// 높은 우선순위가 있으면 stack에 넣고 다음으로
			queue = append(queue, i)
			continue
		}

		// 높은 우선순위가 없으면 i를 실행

		// 누적합 갱신
		for j := priorities[i]; j <= 9; j++ {
			prCount[j]--
		}

		// 실행 순서와 location이 같으면 종료
		if i == location {
			return order
		}

		// 실행 순서 증가
		order++
	}

	return order
}
