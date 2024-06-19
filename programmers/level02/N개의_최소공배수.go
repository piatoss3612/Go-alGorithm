package level02

// 문제: https://school.programmers.co.kr/learn/courses/30/lessons/12953
// 분류: 최소공배수, 유클리드 호제법
func solution(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	lcmValue := lcm(arr[0], arr[1])
	for i := 2; i < len(arr); i++ {
		lcmValue = lcm(lcmValue, arr[i])
	}

	return lcmValue
}

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
