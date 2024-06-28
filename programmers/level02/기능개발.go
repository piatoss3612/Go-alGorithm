package main

// 문제: https://programmers.co.kr/learn/courses/30/lessons/42586
// 분류: 스택/큐
func solution(progresses []int, speeds []int) []int {
	length := len(progresses)       // 작업의 개수
	stack := make([]int, 0, length) // 스택
	answer := []int{}               // 정답
	cnt := 0                        // 함께 배포되는 작업의 개수

	for i := 0; i < length; i++ {
		remain := 100 - progresses[i] // 남은 작업량
		days := remain / speeds[i]    // 작업 완료까지 걸리는 일수
		if remain%speeds[i] != 0 {
			days++
		}

		// 스택이 비어있으면 스택에 추가하고 함께 배포되는 작업의 개수를 1로 초기화
		if len(stack) == 0 {
			stack = append(stack, days)
			cnt++
			continue
		}

		top := stack[len(stack)-1] // 스택의 맨 위에 있는 작업이 완료될 때까지 걸리는 일수

		// 스택의 맨 위에 있는 작업이 완료될 때까지 걸리는 일수보다
		// 현재 작업이 완료될 때까지 걸리는 일수가 작거나 같으면 함께 배포
		if top >= days {
			cnt++
		} else {
			// 가장 위에 있는 작업이 완료될 때까지 걸리는 일수보다 현재 작업이 완료될 때까지 걸리는 일수가 크면
			// 해당 작업을 마무리하고 배포, 함께 배포되는 작업의 개수를 스택에 추가
			// 함께 배포되는 작업의 개수를 1로 초기화
			answer = append(answer, cnt)
			stack = append(stack, days)
			cnt = 1
		}
	}

	// 마지막으로 함께 배포되는 작업의 개수를 추가
	if cnt > 0 {
		answer = append(answer, cnt)
	}

	return answer
}
