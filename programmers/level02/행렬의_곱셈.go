package level02

// 문제: https://school.programmers.co.kr/learn/courses/30/lessons/12949
// 분류: 구현
func solution(arr1 [][]int, arr2 [][]int) [][]int {
	answer := make([][]int, len(arr1))
	for i := range answer {
		answer[i] = make([]int, len(arr2[0]))
		for j := range answer[i] {
			for k := range arr1[i] {
				answer[i][j] += arr1[i][k] * arr2[k][j]
			}
		}
	}
	return answer
}
