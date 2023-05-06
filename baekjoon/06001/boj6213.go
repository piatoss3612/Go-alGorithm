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

	N, Q     int
	A, B     int
	cows     [50001]int
	segments [150000][2]int
)

// 난이도: Gold 1
// 메모리: 8092KB
// 시간: 156ms
// 분류: 세그먼트 트리
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, Q = scanInt(), scanInt()

	for i := 1; i <= N; i++ {
		cows[i] = scanInt()
	}

	NewSegmentTree(1, N, 1)
}

func Solve() {
	for i := 1; i <= Q; i++ {
		A, B = scanInt(), scanInt()
		x, y := Query(A, B, 1, 1, N)
		fmt.Fprintln(writer, x-y)
	}
}

func NewSegmentTree(left, right, node int) (int, int) {
	if left == right {
		segments[node] = [2]int{cows[left], cows[right]}
		return segments[node][0], segments[node][1]
	}

	mid := (left + right) / 2

	lmax, lmin := NewSegmentTree(left, mid, node*2)
	rmax, rmin := NewSegmentTree(mid+1, right, node*2+1)
	segments[node] = [2]int{max(lmax, rmax), min(lmin, rmin)}

	return segments[node][0], segments[node][1]
}

func Query(left, right, node, nodeLeft, nodeRight int) (int, int) {
	if right < nodeLeft || nodeRight < left {
		return -987654321, 987654321
	}

	if left <= nodeLeft && nodeRight <= right {
		return segments[node][0], segments[node][1]
	}

	mid := (nodeLeft + nodeRight) / 2
	lmax, lmin := Query(left, right, node*2, nodeLeft, mid)
	rmax, rmin := Query(left, right, node*2+1, mid+1, nodeRight)

	return max(lmax, rmax), min(lmin, rmin)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
