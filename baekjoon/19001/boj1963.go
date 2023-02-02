package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	isPrime [10000]bool
	T, A, B int
)

// 네 자리 소수를 찾기 위해 에라토스테네스의 체를 사용
func init() {
	for i := 2; i < 10000; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i < 10000; i++ {
		if isPrime[i] {
			for j := i * i; j < 10000; j += i {
				isPrime[j] = false
			}
		}
	}
}

// 난이도: Gold 4
// 메모리: 1200KB
// 시간: 8ms
// 분류: 에라토스테네스의 체, 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	T = scanInt()
	for i := 1; i <= T; i++ {
		Setup()
		Solve()
	}
}

func Setup() {
	A, B = scanInt(), scanInt()
}

func Solve() {
	var checked [10000]bool // 같은 네 자리 소수를 반복해서 큐에 집어넣지 않기 위해 이미 만들어 본 소수인지 체크

	// 너비 우선 탐색
	queue := [][2]int{}
	queue = append(queue, [2]int{A, 0}) // 비밀번호 A를 0번 변경한 상태에서 시작
	checked[A] = true

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		pwd := front[0]
		cnt := front[1]

		// 너비 우선 탐색 종료 조건: B와 일치하는 비밀번호를 생성한 경우
		if pwd == B {
			fmt.Fprintln(writer, cnt)
			return
		}

		splited := split(pwd) // 현재 비밀번호를 네 자리의 배열로 분할

		for i := 0; i < 4; i++ {
			cache := splited[i] // 비밀번호의 i번째 자리 캐싱

			for j := 0; j < 10; j++ {
				// j가 캐시된 수와 동일하거나 비밀번호 최상위 자리의 값을 0으로 변경해야 하는 경우(문제 조건: 1000미만의 비밀번호로는 변경 불가)
				if j == cache || (i == 3 && j == 0) {
					continue
				}
				splited[i] = j // i번째 자리의 수를 j로 변경

				// 새로 변경된 비밀번호의 유효성 검사: 소수이면서 아직 생성한 적이 없는 네 자리수
				if newPwd := join(splited); isPrime[newPwd] && !checked[newPwd] {
					checked[newPwd] = true
					queue = append(queue, [2]int{newPwd, cnt + 1})
				}
			}
			splited[i] = cache // i번째 자리의 수를 초기 상태로 되돌리기
		}
	}

	// 모든 변경 시도를 해보았지만 A를 B로 변경할 수 없는 경우
	fmt.Fprintln(writer, "Impossible")
}

// split: 네자리 수 x를 길이가 4인 배열로 분할
func split(x int) [4]int {
	var result [4]int
	for i := 0; i < 4; i++ {
		result[i] = x % 10
		x /= 10
	}
	return result
}

// join: 길이가 4인 배열을 조합하여 네자리 수를 생성
func join(x [4]int) int {
	result := 0
	digit := 1
	for i := 0; i < 4; i++ {
		result += x[i] * digit
		digit *= 10
	}
	return result
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
