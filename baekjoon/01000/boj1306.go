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

	N, M     int
	lights   [1000001]int
	segments [2200000]int
)

// 난이도: Platinum 5
// 메모리: 31028KB
// 시간: 548ms
// 분류: 세그먼트 트리
// 풀이: 최대값을 저장하는 세그먼트 트리를 만들고 N-2M+2개의 각 구간을 탐색하며 최대값을 출력한다.
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		lights[i] = scanInt()
	}
	NewSegmentTree(1, N, 1)
}

func Solve() {
	for i := M; i <= N-M+1; i++ {
		fmt.Fprintf(writer, "%d ", Query(i-(M-1), i+(M-1), 1, 1, N))
	}
	fmt.Fprintln(writer)
}

func NewSegmentTree(left, right, node int) int {
	if left == right {
		segments[node] = lights[left]
		return segments[node]
	}

	mid := (left + right) / 2
	segments[node] = max(NewSegmentTree(left, mid, node*2),
		NewSegmentTree(mid+1, right, node*2+1))
	return segments[node]
}

func Query(left, right, node, nodeLeft, nodeRight int) int {
	if right < nodeLeft || nodeRight < left {
		return 0
	}

	if left <= nodeLeft && nodeRight <= right {
		return segments[node]
	}

	mid := (nodeLeft + nodeRight) / 2
	return max(Query(left, right, node*2, nodeLeft, mid),
		Query(left, right, node*2+1, mid+1, nodeRight))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
