package main

// 문제: https://school.programmers.co.kr/learn/courses/30/lessons/42839
// 분류: 완전탐색

var (
	notPrime []bool
	checked  []bool
	set      map[int]bool
)

func solution(numbers string) int {
	notPrime = make([]bool, 10000000)
	notPrime[0] = true
	notPrime[1] = true
	for i := 2; i < 10000000; i++ {
		if !notPrime[i] {
			for j := i * i; j < 10000000; j += i {
				notPrime[j] = true
			}
		}
	}

	checked = make([]bool, len(numbers))
	set = make(map[int]bool)

	bruteForce(numbers, 0)

	return len(set)
}

func bruteForce(numbers string, n int) {
	if !notPrime[n] {
		set[n] = true
	}

	for i := 0; i < len(numbers); i++ {
		if checked[i] {
			continue
		}

		checked[i] = true
		bruteForce(numbers, n*10+int(numbers[i]-'0'))
		checked[i] = false
	}
}

func main() {
	println(solution("17"))  // 3
	println(solution("011")) // 2
}
