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
	amount  []int
)

// 최악의 경우: 100000일 동안 매일 10000원을 써야 하는데 1번만 돈을 인출할 수 있을 때, 인출해야 하는 금액
const INF = 1e9

// 난이도: Silver 2
// 메모리: 2268KB
// 시간: 24ms
// 분류: 이분 탐색, 매개 변수 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N, M = scanInt(), scanInt()
	amount = make([]int, N)
	for i := 0; i < N; i++ {
		amount[i] = scanInt()
	}
}

func Solve() {
	l, r := 1, int(INF)
	ans := int(INF) // 인출 금액

	// 인출 금액 이분 탐색
	for l <= r {
		m := (l + r) / 2
		withdrawCount := 0 // 인출 횟수
		balance := 0       // 들고 있는 잔액

		for i := 0; i < N; i++ {
			// 인출 금액이 사용할 금액보다 작은 경우
			// 절대 인출 횟수를 M번 이하로 맞출 수 없으므로 반복문 종료
			if m < amount[i] {
				withdrawCount = INF
				break
			}

			if balance >= amount[i] {
				balance -= amount[i]
			} else {
				balance = m - amount[i]
				withdrawCount++
			}
		}

		if withdrawCount > M {
			// 인출 횟수가 M보다 크다면 한번에 더 큰 금액을 인출해야 한다
			// 따라서 더 큰 금액을 인출하기 위해 이분 탐색 범위 갱신
			l = m + 1
		} else {
			// 인출 횟수가 M보다 작은 경우는
			// 임의로 인출 횟수를 늘려 M으로 맞출 수 있으므로
			// 정확히 M번 인출한 것과 마찬가지라고 간주할 수 있다
			ans = min(ans, m) // 인출 금액의 최솟값 갱신

			// 더 작은 금액을 인출하기 위해 이분 탐색 범위 갱신
			r = m - 1
		}
	}
	fmt.Fprintln(writer, ans)
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
