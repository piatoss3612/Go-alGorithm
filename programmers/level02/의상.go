package main

// 문제: https://school.programmers.co.kr/learn/courses/30/lessons/42578
// 분류: 해시
func solution(clothes [][]string) int {
	clothesMap := make(map[string][]string)

	for _, r := range clothes {
		kind, name := r[1], r[0]
		clothesMap[kind] = append(clothesMap[kind], name)
	}

	answer := 1

	for _, v := range clothesMap {
		answer *= (len(v) + 1) // 의상의 개수 + 의상을 안입는 경우 1
	}

	answer -= 1 // 모든 의상을 안입는 경우 제외

	return answer
}
