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

	T, N    int
	visited [61]bool
	dp      [61][6]int
)

// 난이도: Gold 5
// 메모리: 988KB
// 시간: 4ms
// 시간 복잡도: O(T + N) -> T: 테스트 케이스의 개수, N: BFS를 위한 최대 큐의 크기
// 공간 복잡도: O(N * 6)
// 분류: 그리디 알고리즘, 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	T = scanInt()
	BFS()

	for i := 1; i <= T; i++ {
		Input()
		Solve()
	}
}

func BFS() {
	q := [][6]int{}
	start := [6]int{0, 0, 0, 0, 0, 0}
	q = append(q, start)

	for len(q) > 0 {
		x := q[0]
		q = q[1:]

		if x[0] >= 0 && x[0] <= 60 && !visited[x[0]] {
			visited[x[0]] = true
			dp[x[0]] = x

			q = append(q, [6]int{x[0] - 1, x[1], x[2], x[3], x[4], x[5] + 1})
			q = append(q, [6]int{x[0] + 1, x[1], x[2], x[3], x[4] + 1, x[5]})
			q = append(q, [6]int{x[0] - 10, x[1], x[2], x[3] + 1, x[4], x[5]})
			q = append(q, [6]int{x[0] + 10, x[1], x[2] + 1, x[3], x[4], x[5]})
			q = append(q, [6]int{x[0] + 60, x[1] + 1, x[2], x[3], x[4], x[5]})
		}
	}
}

func Input() {
	N = scanInt()
}

func Solve() {
	q := N / 60
	r := N % 60
	fmt.Fprintf(writer, "%d %d %d %d %d\n", dp[r][1]+q, dp[r][2], dp[r][3], dp[r][4], dp[r][5])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
