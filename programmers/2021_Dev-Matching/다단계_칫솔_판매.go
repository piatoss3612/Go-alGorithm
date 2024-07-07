package main

import "fmt"

// 문제: https://school.programmers.co.kr/learn/courses/30/lessons/77486
// 분류: 구현, 트리
func solution(enroll []string, referral []string, seller []string, amount []int) []int {
	parents := make(map[string]string)
	profits := make(map[string]int)

	for i, e := range enroll {
		parents[e] = referral[i]
		profits[e] = 0
	}

	for i, s := range seller {
		// 10% 제외하고 s의 수익
		total := amount[i] * 100
		fee := total / 10
		profits[s] += total - fee

		current := s

		for {
			// current의 부모가 존재하는지 확인
			parent, ok := parents[current]
			// 부모가 존재하지 않으면 상납금을 먹고 종료
			if !ok {
				profits[current] += fee
				break
			}

			// 자식으로부터 상납금이 1원 미만인 경우 종료
			if fee == 0 {
				break
			}

			// parent의 수익 및 상납금 계산
			total = fee
			fee = total / 10
			profits[parent] += total - fee

			// 부모가 상납금을 납부할 차례
			current = parent
		}
	}

	answer := make([]int, len(enroll))
	for i, s := range enroll {
		answer[i] = profits[s]
	}

	return answer
}

func main() {
	answer := solution(
		[]string{"john", "mary", "edward", "sam", "emily", "jaimie", "tod", "young"},
		[]string{"-", "-", "mary", "edward", "mary", "mary", "jaimie", "edward"},
		[]string{"young", "john", "tod", "emily", "mary"},
		[]int{12, 4, 2, 5, 10},
	)

	fmt.Println(answer)
}
