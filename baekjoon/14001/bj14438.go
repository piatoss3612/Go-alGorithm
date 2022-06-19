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
	n, m    int
	inp     [100001]int
	seg     [270000]int
)

// 메모리: 6908KB
// 시간: 96ms
// 세그먼트 트리를 사용하여 값을 갱신하고 구간의 최솟값을 구하는 문제
// 요즘 문제 푸는게 너무 힘들다...
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n = scanInt()
	for i := 1; i <= n; i++ {
		inp[i] = scanInt()
	}

	SegmentTree(1, n, 1)

	m = scanInt()
	var x, y, z int
	for i := 1; i <= m; i++ {
		x, y, z = scanInt(), scanInt(), scanInt()
		if x == 1 {
			Update(y, z, 1, n, 1)
			inp[y] = z
		} else {
			fmt.Fprintln(writer, Query(y, z, 1, n, 1))
		}
	}
}

func SegmentTree(left, right, node int) int {
	if left == right {
		seg[node] = inp[left]
		return seg[node]
	}

	mid := (left + right) / 2
	seg[node] = min(SegmentTree(left, mid, node*2),
		SegmentTree(mid+1, right, node*2+1))
	return seg[node]
}

func Update(target, newVal, left, right, node int) int {
	if target < left || target > right {
		return seg[node]
	}

	if left == right {
		seg[node] = newVal
		return seg[node]
	}

	mid := (left + right) / 2
	seg[node] = min(Update(target, newVal, left, mid, node*2),
		Update(target, newVal, mid+1, right, node*2+1))
	return seg[node]
}

func Query(left, right, nodeLeft, nodeRight, node int) int {
	if right < nodeLeft || nodeRight < left {
		return 9876543210
	}

	if left <= nodeLeft && nodeRight <= right {
		return seg[node]
	}

	mid := (nodeLeft + nodeRight) / 2
	return min(Query(left, right, nodeLeft, mid, node*2),
		Query(left, right, mid+1, nodeRight, node*2+1))
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
