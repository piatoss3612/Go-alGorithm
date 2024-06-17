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
	A, B    int
	C, D    int
	K       int
)

// 31498번: 장난을 잘 치는 토카 양
// hhttps://www.acmicpc.net/problem/31498
// 난이도: 골드 5
// 메모리: 856 KB
// 시간: 8 ms
// 분류: 이분 탐색, 매개 변수 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	A, B = scanInt(), scanInt()
	C, D = scanInt(), scanInt()
	K = scanInt()
}

func Solve() {
	// 토카는 집에서 A만큼 떨어져 있고 한 번에 B만큼 이동할 수 있다.
	// 토카는 한 번 이동할 때마다 K만큼 이동거리가 줄어든다.
	// 돌돌이는 집과 A+C만큼 떨어져 있고 한 번에 D만큼 이동할 수 있다.

	// K가 0인 경우
	if K == 0 {
		// 토카가 집에 도착하기 위해 필요한 이동횟수
		cnt := func() int {
			if A%B > 0 {
				return A/B + 1
			}
			return A / B
		}()

		// 돌돌이가 토카를 따라잡을 수 있는지 확인
		if doldolReachhHome(cnt) {
			fmt.Fprintln(writer, -1)
		} else {
			fmt.Fprintln(writer, cnt)
		}
		return
	}

	// 이동거리가 줄어듬에 따른 토카의 최대 이동가능 횟수
	maxMove := func() int {
		if B/K > 0 {
			return B/K + 1
		}
		return B / K
	}()

	// 최대로 이동했을 때 집에 도착하는지 확인
	if !tokaReachHome(maxMove) {
		fmt.Fprintln(writer, -1)
		return
	}

	// 이진탐색으로 최소 이동횟수 찾기
	l, r := 1, maxMove
	for l <= r {
		mid := (l + r) / 2
		if tokaReachHome(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	// 최소로 이동하여 집에 도착했음에도 돌돌이가 토카를 따라잡을 수 있는지 확인
	if doldolReachhHome(l) {
		fmt.Fprintln(writer, -1)
		return
	}

	fmt.Fprintln(writer, l)
}

func tokaMove(n int) int {
	return A - B*n + n*(n-1)*K/2
}

func tokaReachHome(n int) bool {
	return tokaMove(n) <= 0
}

func doldolMove(n int) int {
	return A + C - D*n
}

func doldolReachhHome(n int) bool {
	return doldolMove(n) <= 0
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
