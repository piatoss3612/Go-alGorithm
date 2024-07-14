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
	n, k    int
)

// 20444번: 색종이와 가위
// hhttps://www.acmicpc.net/problem/20444
// 난이도: 골드 5
// 메모리: 860 KB
// 시간: 4 ms
// 분류: 이분 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	n, k = scanInt(), scanInt()
}

func Solve() {
	// 색종이를 n번 잘라서 k개의 조각을 만들 수 있는지 확인
	// 색종이를 a번 가로로 자르고 b번 세로로 자르면 (a+1)*(b+1)개의 조각이 생김
	// a, b는 0부터 n까지 가능

	// 가로 또는 세로로 자르는 횟수의 upper bound를 이분 탐색으로 찾음
	l, r := 0, n

	for l <= r {
		mid := (l + r) / 2 // 가로로 자르는 횟수

		pieces := (mid + 1) * (n - mid + 1) // 가로로 mid번 자르고 세로로 n-mid번 자르면 만들어지는 조각 수

		// k개의 조각을 만들 수 있으면 YES 출력 후 종료
		if pieces == k {
			fmt.Fprintln(writer, "YES")
			return
		}

		// k개보다 적게 만들어지면 가로로 자르는 횟수를 늘려야 함
		// 왜? n이 10이라고 했을 때, n의 l을 5라고 가정했을 때 6*6=36개의 조각이 만들어짐
		// 그런데 l이 10의 절반인 5보다 작아지거나 커지면 조각의 수가 줄어듬 (5*6=30, 4*7=28)
		// 따라서 조각의 수가 k보다 크면 mid를 줄이는 방향으로 이분 탐색을 진행해야 함 (늘려도 되는데 일관성을 위해 줄임)
		// 조각의 수가 k보다 작으면 mid를 늘리는 방향으로 이분 탐색을 진행
		if pieces < k {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	// k개의 조각을 만들 수 없으면 NO 출력
	fmt.Fprintln(writer, "NO")
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	return mustParseInt(scanString())
}
