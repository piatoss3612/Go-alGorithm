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
	N, M    int
	weight  [31]int
	bead    [8]int
	dp      [15001]bool // 구슬의 무게를 확인할 수 있는지 여부를 저장하는 배열, 추의 개수 * 추의 무게의 최댓값 = 15000
)

// 난이도: Gold 3
// 메모리: 924KB
// 시간: 4ms
// 분류: 다이나믹 프로그래밍, 배낭 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N = scanInt()
	for i := 1; i <= N; i++ {
		weight[i] = scanInt()
	}
	M = scanInt()
	for i := 1; i <= M; i++ {
		bead[i] = scanInt()
	}
}

func Solve() {
	dp[0] = true // 구슬의 무게가 0인 경우(저울에 구슬이 올라가있지 않은 경우)는 항상 확인이 가능하다

	for w := 1; w <= N; w++ {
		// 1. 구슬의 반대편에 w번째 추 올리기
		// 더하는 과정의 중복을 피하기 위해 내림차순으로 진행
		for s := 15000; s >= 0; s-- {
			// 무게가 s인 구슬을 확인할 수 있는 경우
			if dp[s] {
				if s+weight[w] <= 15000 {
					// 무게가 s + w번째 추의 무게인 구슬도 확인할 수 있어야 한다
					dp[s+weight[w]] = true
				}

			}
		}

		// 2. 구슬이 있는 쪽에 w번째 추 올리기
		// 빼는 과정의 중복을 피하기 위해 오름차순으로 진행
		for s := 0; s <= 15000; s++ {
			// 무게가 s인 구슬을 확인할 수 있는 경우
			if dp[s] {
				// 무게가 abs(s-weight[w])인 구슬도 확인할 수 있어야 한다
				dp[abs(s-weight[w])] = true
			}
		}
	}

	// 구슬의 무게를 확인할 수 있는지 여부 출력
	for i := 1; i <= M; i++ {
		if bead[i] <= 15000 && dp[bead[i]] {
			fmt.Fprint(writer, "Y ")
		} else {
			fmt.Fprint(writer, "N ")
		}
	}
	fmt.Fprintln(writer)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
