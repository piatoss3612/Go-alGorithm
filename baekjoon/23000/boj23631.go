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

	T, N, K int
)

// 난이도: Silver 1
// 메모리: 5460KB
// 시간: 76ms
// 분류: 이분 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()

}

func Setup() {
	T = scanInt()

	for i := 0; i < T; i++ {
		N, K = scanInt(), scanInt()
		Solve()
	}
}

func Solve() {
	l, r := 1, 10000 // (x * (x+1)) / 2 >= 10000000을 만족하는 적당한 x를 r로 사용

	for l <= r {
		mid := (l + r) / 2
		// mid번 왕복했을 때의 이동거리가 N-1보다 크거나 같으면
		if possible(mid) {
			r = mid - 1 // r을 줄여서 더 작은 mid를 찾아본다
		} else {
			l = mid + 1 // l을 늘려서 더 큰 mid를 찾아본다
		}
	}

	var dir string 			   // l번 왕복했을 때의 방향

	pos := K * ((l + 1) / 2)       // l번 왕복했을 때의 위치
	moved := K * (l * (l + 1)) / 2 // l번 왕복했을 때의 이동거리
	diff := moved - (N - 1)        // l번 왕복했을 때의 이동거리와 정확히 N-1만큼 이동했을 때의 이동거리의 차이

	if l%2 == 0 {
		dir = "L"
		pos = -pos // 왼쪽으로 이동했으므로 pos는 음수
		pos += diff // diff만큼 더 이동했으므로 pos에 diff를 더해준다
	} else {
		dir = "R"
		pos -= diff // diff만큼 더 이동했으므로 pos에서 diff를 빼준다
	}

	fmt.Fprintln(writer, pos, dir)
}

func possible(mid int) bool {
	return (mid*(mid+1))*K/2 >= N // mid번 왕복했을 때의 이동거리가 N보다 크거나 같으면 true
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
