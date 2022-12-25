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
	N, K, C int
	chef    []int // 각 요리사의 요리 시간
	ans     = MAX
)

const MAX = 1000000000000 // 1 <= K * Ai <= 1000000000000

// 난이도: Gold 5
// 메모리: 916KB
// 시간: 480ms
// 분류: 브루트포스 알고리즘, 이분 탐색, 백트래킹, 매개 변수 탐색

// 요리사를 격려한다는게 요리를 만들 때마다 한 번씩 할 수 있는 것인지
// 요리를 시작하기 전에 한 번에 격려할 수 있는 것인지 모호한 부분이 있다.
// 결과적으로는 한 번에 격려하는 방법이 맞았다.
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N, K, C = scanInt(), scanInt(), scanInt()
	chef = make([]int, N)
	for i := 0; i < N; i++ {
		chef[i] = scanInt()
	}
}

func Solve() {
	BackTracking(0)
	fmt.Fprintln(writer, ans)
}

// 백트래킹 + 브루트포스: N명의 요리사에게 C번 나누어 격려할 수 있는 모든 경우의 수 고려
func BackTracking(encouraged int) {
	// encouraged 횟수 만큼 격려를 한 상황에서 K개의 음식 조리가 완료되는 최소 시간 찾기
	BinarySearch()

	// 기저사례: C번 격려를 마친 경우
	if encouraged == C {
		return
	}

	// 백트래킹
	for i := 0; i < N; i++ {
		// i번 요리사를 격려할 수 있는 경우
		if chef[i] > 1 {
			chef[i]--
			// i번 요리사를 격려했을 때의 최소 시간 찾기
			BackTracking(encouraged + 1)
			chef[i]++
		}
	}
}

func BinarySearch() {
	l, r := 1, MAX
	for l <= r {
		mid := (l + r) / 2 // 매개 변수: K개의 음식 조리가 완료되는데 걸리는 시간

		dishes := 0 // 완성된 요리의 개수
		for _, c := range chef {
			dishes += mid / c
			if dishes >= K {
				break
			}
		}

		if dishes >= K {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	ans = min(ans, l) // l은 K개의 음식 조리가 완료되는데 걸리는 시간의 최솟값
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
