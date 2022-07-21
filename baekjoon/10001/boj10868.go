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
	inp     [100001]int
	minTree [262145]int // 세그먼트 트리의 크키: 100000보다 큰 2의 거듭제곱  131072*2 + 1
	n, m    int
)

// 메모리: 8252KB
// 시간: 124ms
// 2357번 문제와 풀이 동일
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()
	for i := 1; i <= n; i++ {
		inp[i] = scanInt()
	}

	minSegment(1, n, 1)

	var a, b int
	for i := 1; i <= m; i++ {
		a, b = scanInt(), scanInt()
		fmt.Fprintln(writer, query(a, b, 1, 1, n))
	}
}

func query(left, right, node, nodeLeft, nodeRight int) int {
	if right < nodeLeft || left > nodeRight {
		return 1000000001
	}

	if left <= nodeLeft && nodeRight <= right {
		return minTree[node]
	}

	mid := (nodeLeft + nodeRight) / 2
	return min(query(left, right, node*2, nodeLeft, mid), query(left, right, node*2+1, mid+1, nodeRight))
}

func minSegment(left, right, node int) int {
	if left == right {
		minTree[node] = inp[left]
		return minTree[node]
	}

	mid := (left + right) / 2
	minTree[node] = min(minSegment(left, mid, node*2), minSegment(mid+1, right, node*2+1))
	return minTree[node]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
