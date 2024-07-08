package main

import "fmt"

// 문제: https://school.programmers.co.kr/learn/courses/30/lessons/118667
// 분류: 두 포인터
func solution(queue1 []int, queue2 []int) int {
	q1Sum, q2Sum := 0, 0 // 큐1, 큐2의 합

	for i := 0; i < len(queue1); i++ {
		q1Sum += queue1[i]
		q2Sum += queue2[i]
	}

	target := (q1Sum + q2Sum) / 2 // 큐1의 합과 큐2의 합이 동일한 경우의 합

	entire := append(queue1, queue2...) // 큐1과 큐2를 합친 배열

	// queue1을 기준으로 투 포인터 알고리즘을 사용하여 target을 만족하는 경우를 찾는다.
	// queue1의 합이 target을 만족하면, queue2의 합은 자동으로 target을 만족하게 된다.
	sum := q1Sum             // 큐1의 합
	l, r := 0, len(queue1)-1 // 투 포인터 (queue1의 시작과 끝)
	cnt := 0                 // 각 큐에서 원소를 추출하여 상대 큐에 넣는 횟수

	// l과 r이 entire 배열의 범위 내에 있을 때까지 반복
	for l < len(entire) && r < len(entire) {
		// 큐1의 합이 target을 만족하면 종료
		if sum == target {
			return cnt
		}

		if sum < target && r+1 < len(entire) { // 큐1의 합이 target보다 작으면, 큐2에서 원소를 하나 추출하여 큐1에 넣는다.
			r++
			sum += entire[r]
		} else { // 큐1의 합이 target보다 크면, 큐1에서 원소를 하나 추출하여 큐2에 넣는다.
			sum -= entire[l]
			l++
		}

		cnt++ // 횟수 증가
	}

	return -1
}

func main() {
	answer := solution([]int{3, 2, 7, 2}, []int{4, 6, 5, 1})
	fmt.Println(answer) // Expect 2
	answer = solution([]int{1, 2, 1, 2}, []int{1, 10, 1, 2})
	fmt.Println(answer) // Expect 7
	answer = solution([]int{1, 1}, []int{1, 5})
	fmt.Println(answer) // Expect -1
}
