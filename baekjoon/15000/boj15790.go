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
	N, M, K int
	grooves []int
)

// 난이도: Gold 3
// 메모리: 924KB
// 시간: 316ms
// 분류: 이분 탐색, 매개 변수 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, M, K = scanInt(), scanInt(), scanInt()
	grooves = make([]int, M)
	for i := 0; i < M; i++ {
		grooves[i] = scanInt()
	}
	// 홈의 위치는 오름차순으로 주어지므로 정렬 불필요
}

func Solve() {
	l, r := 1, 100000
	for l <= r {
		m := (l + r) / 2
		// 길이가 m인 고무줄을 K개를 만들어 활을 만들 수 있는 경우
		if CanMakeBow(m) {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	if r == 0 {
		fmt.Fprintln(writer, -1) // r이 0인 경우는 활을 만들 수 없다
	} else {
		fmt.Fprintln(writer, r) // 최댓값(upper bound 출력)
	}
}

func CanMakeBow(expected int) bool {
	// 어느 홈에서 시작해야 최적해를 구할 수 있는지 알 수 없기 때문에
	// 모든 홈에서 각각 시작하여 길이가 expected인 고무줄을 K개 이상 만들 수 있는지 확인해야 한다
	for i := 0; i < M; i++ {
		left := N          // 남은 고무줄의 길이
		from := grooves[i] // 현재 위치
		cnt := 0           // 길이가 expected 이상인 고무줄의 개수

		for j := 1; j < M; j++ {
			to := grooves[(i+j)%M]        // 다음 홈의 위치
			length := (to - from + N) % N // 다음 홈의 위치에서 고무줄을 잘랐을 때, 잘라낸 고무줄의 길이

			// 잘라낸 고무줄의 길이가 expected 이상인 경우
			if length >= expected {
				left -= length
				from = to
				cnt += 1
			}
		}

		// 남은 고무줄의 길이가 expected 이상인 경우
		if left >= expected {
			cnt += 1
		}

		// 길이가 expected 이상인 고무줄의 개수가 K개 이상인 경우
		if cnt >= K {
			return true
		}
	}

	// 활을 만들 수 없는 경우
	return false
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
